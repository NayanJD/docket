package models

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getMysqlDsn() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database_name := os.Getenv("DB_DATABASE")

	return fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?parseTime=true",
		username,
		password,
		host,
		port,
		database_name,
	)
}

func GetDB() *gorm.DB {
	return DB
}

func ConnectDatabase() {

	mysqlDsn := getMysqlDsn()

	database, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to mysql")
	}

	log.Info().Msg("Successfully connected to mysql")
	log.Info().Msg("Test restart")

	err = database.AutoMigrate(&User{})
	database.AutoMigrate(&ClientStoreItem{})

	if err != nil {
		log.Error().Err(err)
	}

	DB = database
}
