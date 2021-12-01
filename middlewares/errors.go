package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		c.Next()

		if len(c.Errors) > 0 {
			log.Error().Msg(c.Errors.String())
		}
	}

}

func RecoveryFunc(c *gin.Context, recovered interface{}) {
	if err, ok := recovered.(string); ok {
		log.Error().Msg(err)
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
}