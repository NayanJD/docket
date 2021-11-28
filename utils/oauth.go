package utils

import (
	"fmt"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"

	oauthMysqlStore "github.com/go-oauth2/mysql/v4"

	"nayanjd/docket/models"
)

func SetupOauth() *server.Server {
	manager := manage.NewDefaultManager()
	
	// use mysql token store
	mysqlStore := oauthMysqlStore.NewDefaultStore(
		oauthMysqlStore.NewConfig(GetMysqlDsn()),
	)
	clientStore, _ := models.NewClientStore(models.DB)

	manager.MapTokenStorage(mysqlStore)
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