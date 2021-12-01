package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){

		//To catch unknown panics anywhere
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			}
		}()

		c.Next()

		if len(c.Errors) > 0 {
			log.Error().Msg(c.Errors.String())
		}
	}

}