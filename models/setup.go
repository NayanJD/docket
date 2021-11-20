package models

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(postgres.Open("postgres://go:password@localhost:5432/docket_local"), &gorm.Config{})

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