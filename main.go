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
	"nayanjd/docket/utils"
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

	userController := controllers.UserController{}

	userEndpoints := r.Group("user")
	{
		userEndpoints.POST(
			"/register",
			middlewares.JSONValidationMiddleware(controllers.UserInputForm{}),
			userController.Register,
		)

		userEndpoints.GET("", oauthController.TokenMiddleware(), userController.GetUsers)
		userEndpoints.GET("/:id", oauthController.TokenMiddleware(), userController.GetUser)
	}

	taskEndpoints := r.Group("task")
	{
		taskController := controllers.TaskController{}

		taskEndpoints.POST(
			"",
			oauthController.TokenMiddleware(),
			middlewares.JSONValidationMiddleware(controllers.TaskInputForm{}),
			taskController.Create,
		)

		taskEndpoints.GET(
			"",
			oauthController.TokenMiddleware(),
			taskController.GetUserTasks,
		)

		taskEndpoints.PUT(
			"/:id",
			oauthController.TokenMiddleware(),
			middlewares.JSONValidationMiddleware(controllers.TaskInputForm{}),
			taskController.UpdateUserTask,
		)

		taskEndpoints.PATCH(
			"/:id",
			oauthController.TokenMiddleware(),
			middlewares.JSONValidationMiddleware(controllers.PatchTaskInputForm{}),
			taskController.UpdateUserTask,
		)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.NoRoute(func(c *gin.Context) {
		utils.AbortWithGenericJson(c, nil, &utils.PathNotFoundError)
	})

	log.Printf("Server stopped, err: %v", r.Run(":8000"))
}
