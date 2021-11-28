package controllers

import (
	"nayanjd/docket/utils"
	"net/http"

	"github.com/gin-gonic/gin"
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
	_, err := Srv.ValidationBearerToken(c.Request)
	
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are unauthorised to view this resource",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Test resource"})
	}

}

