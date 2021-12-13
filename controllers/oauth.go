package controllers

import (
	"fmt"
	"nayanjd/docket/models"
	"nayanjd/docket/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// var Srv = utils.SetupOauth()

type OauthController struct{}

type OauthTokenBody struct {
	Client_id     string `json:"client_id"`
	Client_secret string `json:"client_secret"`
	Scope         string `json:"scope"`
	Grant_type    string `json:"grant_type"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}

type OauthTokenData struct {
	Access_token  string `json:"access_token"`
	Expires_in    int    `json:"expires_in"`
	Refresh_token string `json:"refresh_token"`
	Scope         string `json:"scope"`
	Token_type    string `json:"token_type"`
}

type OauthTokenResponse struct {
	utils.GenericResponseBody
	Data OauthTokenData `json:"data"`
}

// TokenHandler Oauth token
// @Summary	Get Oauth bearer token
// @Accept	mpfd
// @Produce json
// @Param	grants	body	OauthTokenBody	true	"Create token"
// @Success	200		{object}	OauthTokenResponse	"Success"
// @Router	/oauth/token	[post]
func (ctrl OauthController) TokenHandler(c *gin.Context) {
	err := utils.GetSrv().HandleTokenRequest(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong password",
		})
	}
}

func (ctrl OauthController) AuthorizeHandler(c *gin.Context) {
	err := utils.GetSrv().HandleAuthorizeRequest(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong password",
		})
	}
}

func (ctrl OauthController) TestHandler(c *gin.Context) {

	user := c.MustGet(gin.AuthUserKey).(models.User)

	log.Info().Msg(fmt.Sprintf("Got user as %v", user))

	utils.AbortWithGenericJson(
		c,
		utils.CreateOKResponse(
			map[string]string{"message": "Test resource success"},
			nil,
		),
		nil,
	)
}

func (ctrl OauthController) TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenInfo, err := utils.GetSrv().
			ValidationBearerToken(c.Request)

		if err != nil {
			utils.AbortWithGenericJson(
				c,
				nil,
				&utils.UnauthorisedError,
			)
			return
		} else {
			user := models.User{}

			if err = models.GetDB().First(&user, "id = ?", tokenInfo.GetUserID()).Error; err != nil {
				utils.AbortWithGenericJson(c, nil, &utils.UnauthorisedError)
				return
			} else {
				c.Set(gin.AuthUserKey, user)
				c.Next()
				return
			}
		}
	}
}
