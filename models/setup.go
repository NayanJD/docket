package models

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database_name := os.Getenv("DB_DATABASE")

	mysqlDsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", username, password, host, port, database_name)

	database, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to mysql")
	}

	log.Info().Msg("Successfully connected to mysql")
	log.Info().Msg("Test restart")

	err = database.AutoMigrate(&User{})
	
	if err != nil {
		log.Error().Err(err)
	}

	database.Migrator().AlterColumn(&User{}, "First_name")

	DB = database
}