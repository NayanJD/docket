package models

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("go:password@tcp(127.0.0.1:3306)/docket_local"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to postgres")
	}

	log.Info().Msg("Successfully connected to postgres")

	err = database.AutoMigrate(&User{})
	
	if err != nil {
		log.Error().Err(err)
	}

	database.Migrator().AlterColumn(&User{}, "First_name")

	DB = database
}