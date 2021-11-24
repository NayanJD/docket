package main

import (
	"net/http"
	"os"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"nayanjd/docket/models"
	"nayanjd/docket/utils"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Warn().Msg("Error loading .env")
	}
	r := gin.New()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	r.Use(ginzerolog.Logger("gin"))

	models.ConnectDatabase()

	
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to docket",
		})
	})

	srv := utils.SetupOauth()
	
	oauthEndpoints := r.Group("oauth")
	{
		oauthEndpoints.POST("/token", func(c *gin.Context) {
			err := srv.HandleTokenRequest(c.Writer, c.Request)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Wrong password",
				})
			}
		})

		oauthEndpoints.POST("/authorize", func(c *gin.Context) {
			err := srv.HandleAuthorizeRequest(c.Writer, c.Request)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Wrong password",
				})
			}
		})
	}

	log.Printf("Server stopped, err: %v", r.Run(":8000"))
}