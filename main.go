package main

import (
	"os"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"nayanjd/docket/controllers"
	_ "nayanjd/docket/docs"
	"nayanjd/docket/middlewares"
	"nayanjd/docket/models"
)

// Swagger Annotations
// @title           Docket API
// @version         1.0
// @description     This is the API for the docket app

// @contact.name   Nayan Das
// @contact.email  dastms@gmail.com

// @host      localhost:8080
func main() {
	err := godotenv.Load()

	if err != nil {
		log.Warn().Msg("Error loading .env")
	}
	r := gin.New()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	r.Use(ginzerolog.Logger("gin"))
	r.Use(middlewares.ErrorMiddleware())
	r.Use(gin.CustomRecovery(middlewares.RecoveryFunc))

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

		oauthEndpoints.GET("/test", oauthController.TokenMiddleware(), oauthController.TestHandler)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Server stopped, err: %v", r.Run(":8000"))
}
