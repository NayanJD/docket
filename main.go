package main

import (
	"os"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"nayanjd/docket/models"
)

func main() {
	r := gin.New()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	r.Use(ginzerolog.Logger("gin"))

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to docket",
		})
	})
	log.Printf("Server stopped, err: %v", r.Run(":8000"))
}