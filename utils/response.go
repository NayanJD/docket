package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Http_code	int
	Body		*gin.H
	Meta		*gin.H
}

type GenericResponseBody struct {
	Data		interface{}	`json:"data"`
	Errors		[]APIError	`json:"errors"`
	IsSuccess	bool		`json:"is_success"`
	Meta		interface{}	`json:"meta"`
}

func IsStatusSuccess(code int) bool {
	return code >= http.StatusOK && code < http.StatusMultipleChoices 
}

func CreateResponse(code int, body *gin.H, meta *gin.H) *Response {
	return &Response{Http_code: code, Body: body, Meta: meta}
}

func CreateOKResponse(body *gin.H, meta *gin.H) *Response {
	return CreateResponse(http.StatusOK, body, meta)
}

func AbortWithGenericJson(c *gin.Context, r *Response, err *APIError) {

	if r != nil {
		body := gin.H{}
		if r.Body != nil {
			body = *r.Body
		}

		meta := gin.H{}
		if r.Meta != nil {
			meta = *r.Meta
		}

		c.AbortWithStatusJSON(r.Http_code, gin.H{
			"data": body,
			"error": nil,
			"isSuccess": IsStatusSuccess(r.Http_code),
			"meta": meta,
		})
		return
	} 

	if err != nil {
		c.AbortWithStatusJSON(err.Http_code, gin.H{
			"data": nil,
			"error": err,
			"isSuccess": IsStatusSuccess(err.Http_code),
			"meta": gin.H{},
		})
		return
	}
}