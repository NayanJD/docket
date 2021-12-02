package controllers

import (
	"fmt"
	"nayanjd/docket/models"
	"nayanjd/docket/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var Srv = utils.SetupOauth()

type OauthController struct{}

func (ctrl OauthController) TokenHandler(c *gin.Context) {
	err := Srv.HandleTokenRequest(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong password",
		})
	}
}

func (ctrl OauthController) AuthorizeHandler(c *gin.Context) {
	err := Srv.HandleAuthorizeRequest(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong password",
		})
	}
}

func (ctrl OauthController) TestHandler(c *gin.Context) {

	user := c.MustGet(gin.AuthUserKey).(models.User)

	log.Info().Msg(fmt.Sprintf("Got user as %v", user))

	utils.AbortWithGenericJson(c, http.StatusOK, gin.H{"message": "Test resource success"})
}

func (ctrl OauthController) TokenMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenInfo, err := Srv.ValidationBearerToken(c.Request)
		
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are unauthorised to view this resource",
			})
			return
		} else {
			user := models.User{}

			if err = models.GetDB().First(&user,"id = ?", tokenInfo.GetUserID()).Error; err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "You are unauthorised to view this resource",
				})
				return
			} else {
				c.Set(gin.AuthUserKey, user)
				c.Next()
				return
			}
		}
	}
}

