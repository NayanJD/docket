package middlewares

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

func JSONValidationMiddleware(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		modelValue := reflect.ValueOf(model)

		typ := modelValue.Type()

		obj := reflect.New(typ).Interface()

		if err := c.ShouldBindJSON(obj); err != nil {
			c.Error(err).SetType(gin.ErrorTypeBind)
			c.Abort()
		} else {
			c.Next()
		}

	}
}
