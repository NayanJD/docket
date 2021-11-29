package main

import (
	"flag"
	"fmt"

	"nayanjd/docket/models"

	"github.com/joho/godotenv"
)



func CreateSuperuser(flagSet *flag.FlagSet, commands []string) {
	firstName := flagSet.String("first-name", "first_name", "First name of user")
	lastName := flagSet.String("last-name", "last_name", "Last name of user")
	username := flagSet.String("username", "username@user.com", "Username (email) of user")
	password := flagSet.String("password", "password", "Password of account")
	isSuperuser := true
	isStaff := false
	flagSet.Parse(commands)

	newUser := models.User{
		First_name: firstName,
		Last_name: lastName,
		Username: username,
		Password: password,
		Is_superuser: &isSuperuser,
		Is_staff: &isStaff,
	}

	godotenv.Load()
	models.ConnectDatabase()

	models.GetDB().Create(&newUser)

	fmt.Println("New superuser created")
}