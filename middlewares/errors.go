package middlewares

import (
	"nayanjd/docket/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				switch err.Type {
				case gin.ErrorTypeBind:
					errs := err.Err.(validator.ValidationErrors)
					list := []string{}
					for _, err := range errs {
						list = append(list, utils.ValidationErrorToText(err))
					}

					utils.AbortWithGenericJson(c, nil, utils.CreateUnprocessableEntityError(list))
				case utils.ErrorTypeDB:
					log.Debug().Msg("DB error occured")
					log.Error().Msg(err.Err.Error())
					utils.AbortWithGenericJson(c, nil, utils.CreateDbError(err.Err))
				default:
					log.Error().Msg("Unknown error occurred")
				}
			}

		}

		if !c.Writer.Written() {
			utils.AbortWithGenericJson(c, nil, &utils.InternalServerError)
		}
	}
}

func RecoveryFunc(c *gin.Context, recovered interface{}) {
	if err, ok := recovered.(string); ok {
		log.Error().Msg(err)
	}

	utils.AbortWithGenericJson(c, nil, &utils.InternalServerError)
}
