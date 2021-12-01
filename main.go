package main

import (
	"os"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"nayanjd/docket/controllers"
	"nayanjd/docket/middlewares"
	"nayanjd/docket/models"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Warn().Msg("Error loading .env")
	}
	r := gin.New()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	r.Use(ginzerolog.Logger("gin"))
	r.Use(middlewares.ErrorMiddleware())

	models.ConnectDatabase()

	
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to docket",
		})
	})

	oauthController := controllers.OauthController{}
	
	oauthEndpoints := r.Group("oauth")
	{
		oauthEndpoints.POST("/token", oauthController.TokenHandler)

		oauthEndpoints.POST("/authorize", oauthController.AuthorizeHandler)

		oauthEndpoints.GET("/test", oauthController.TestHandler)
	}
	
	log.Printf("Server stopped, err: %v", r.Run(":8000"))
}