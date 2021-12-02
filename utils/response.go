package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsStatusSuccess(code int) bool {
	return code >= http.StatusOK && code < http.StatusMultipleChoices 
}

func AbortWithGenericJson(c *gin.Context, code int, obj interface{}) {
	isSuccess := IsStatusSuccess(code)

	if isSuccess {
		c.AbortWithStatusJSON(code, gin.H{
			"data": obj,
			"errors": nil,
			"isSuccess": true,
			"meta": gin.H{},
		})
		return
	} else {
		c.AbortWithStatusJSON(code, gin.H{
			"data": nil,
			"errors": obj,
			"isSuccess": false,
			"meta": gin.H{},
		})
		return
	}
}