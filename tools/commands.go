package main

import (
	"flag"
	"fmt"
	"os"
)



func main() {
	oauthClientCmd := flag.NewFlagSet("oauthClient", flag.ExitOnError)
	superuserCmd := flag.NewFlagSet("superuser", flag.ExitOnError)

	

	flag.Parse()

	if len(os.Args) < 2 {
        fmt.Println("expected 'oauthClient' or 'superuser' subcommands")
		return
    }

	switch os.Args[1] {
		case "oauthClient":
			CreateOauthClient(oauthClientCmd, os.Args[2:])
		case "superuser":
			CreateSuperuser(superuserCmd, os.Args[2:])
		default:
			fmt.Println("expected 'oauthClient' or 'superuser' subcommands")
	}
}