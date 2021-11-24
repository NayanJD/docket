package utils

import (
	"fmt"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	oauthModels "github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/rs/zerolog/log"

	"github.com/go-oauth2/mysql/v4"
)

func SetupOauth() *server.Server {
	manager := manage.NewDefaultManager()
	
	// use mysql token store
	mysqlStore := mysql.NewDefaultStore(
		mysql.NewConfig(GetMysqlDsn()),
	)

	manager.MapTokenStorage(mysqlStore)
	// manager.MapClientStorage(clientStore)

	// client memory store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &oauthModels.Client{
		ID:     "000000",
		Secret: "999999",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	// srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		if username == "test" && password == "test" {
			userID = "test"
		}
		return userID, nil
	})

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Error().Msg(fmt.Sprintf("Internal Error: %v", err.Error()))
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Error().Msg(fmt.Sprintf("Response Error: %v", re.Error.Error()))
	})

	return srv
}