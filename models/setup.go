package models

import (
	"fmt"
	"os"
	"strconv"

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
		log.Warn().Msg("Failed to connect to mysql")
	} else {
		log.Info().Msg("Successfully connected to mysql")
	}

	hasMigrationRun := false
	shouldRunMigrationStr, envOk := os.LookupEnv("RUN_MIGRATIONS")

	if envOk {
		shouldRunMigrations, err := strconv.ParseBool(shouldRunMigrationStr)

		if err == nil && shouldRunMigrations {
			log.Info().Msg("Running migrations")

			err = database.AutoMigrate(&User{})
			database.AutoMigrate(&ClientStoreItem{})

			if err != nil {
				log.Error().Err(err)
			}

			hasMigrationRun = true
		}
	}

	if !hasMigrationRun {
		log.Warn().
			Msg("Migrations has not been run. Set RUN_MIGRATIONS env to true to run them")
	}

	DB = database
}
