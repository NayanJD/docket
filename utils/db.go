package utils

import (
	"fmt"
	"os"
)

func GetMysqlDsn() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database_name := os.Getenv("DB_DATABASE")

	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", username, password, host, port, database_name)
}
