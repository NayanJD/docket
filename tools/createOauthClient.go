package main

import (
	"flag"
	"fmt"

	"nayanjd/docket/models"

	"github.com/joho/godotenv"
)



func main() {
	name := flag.String("name", "client", "Client name for the new client")

	flag.Parse()

	newClient := models.ClientStoreItem{
		Name: *name,
	}

	godotenv.Load()
	models.ConnectDatabase()

	models.GetDB().Create(&newClient)

	fmt.Println("client id: ", newClient.ID)
	fmt.Println("client secret: ", newClient.Secret)
}