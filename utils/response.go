package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Http_code int
	Body      *gin.H
	Meta      *gin.H
}

type GenericResponseBody struct {
	Data      interface{} `json:"data"`
	Errors    []APIError  `json:"errors"`
	IsSuccess bool        `json:"is_success"`
	Meta      interface{} `json:"meta"`
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
			"data":      body,
			"error":     nil,
			"isSuccess": IsStatusSuccess(r.Http_code),
			"meta":      meta,
		})
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(err.Http_code, gin.H{
			"data":      nil,
			"error":     err,
			"isSuccess": IsStatusSuccess(err.Http_code),
			"meta":      gin.H{},
		})
		return
	}
}

func ValidationErrorToText(e validator.FieldError) string {
	lowerCaseField := strings.ToLower(e.Field())
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", lowerCaseField)
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", lowerCaseField, e.Param())
	case "min":
		return fmt.Sprintf("%s must be longer than %s", lowerCaseField, e.Param())
	case "email":
		return fmt.Sprintf("Invalid email format")
	case "len":
		return fmt.Sprintf("%s must be %s characters long", lowerCaseField, e.Param())
	}
	return fmt.Sprintf("%s is not valid", lowerCaseField)
}
