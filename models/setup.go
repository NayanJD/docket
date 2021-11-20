package models

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
		  SlowThreshold:              time.Second,   // Slow SQL threshold
		  LogLevel:                   logger.Info, // Log level
		  IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
		  Colorful:                  false,          // Disable color
		},
	  )
	  
	database, err := gorm.Open(postgres.Open("postgres://go:password@localhost:5432/docket_local"), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("Failed to connect to postgres")
	}

	// log.Info().Msg("Successfully connected to postgres")

	err = database.Debug().AutoMigrate(&User{})
	
	if err != nil {
		// log.Error().Err(err)
	}

	database.Migrator().AlterColumn(&User{}, "first_name")
	database.Migrator().AlterColumn(&User{}, "email")

	DB = database
}