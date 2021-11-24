package models

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"nayanjd/docket/utils"
)

var DB *gorm.DB

func ConnectDatabase() {

	mysqlDsn := utils.GetMysqlDsn()

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