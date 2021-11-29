package main

import (
	"flag"
	"fmt"

	"nayanjd/docket/models"

	"github.com/joho/godotenv"
)



func CreateOauthClient(flagSet *flag.FlagSet, commands []string) {
	name := flagSet.String("name", "client", "Client name for the new client")

	flagSet.Parse(commands)

	newClient := models.ClientStoreItem{
		Name: *name,
	}

	godotenv.Load()
	models.ConnectDatabase()

	models.GetDB().Create(&newClient)

	fmt.Println("client id: ", newClient.ID)
	fmt.Println("client secret: ", newClient.Secret)
}