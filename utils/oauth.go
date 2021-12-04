package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"

	oauthMysqlStore "github.com/go-oauth2/mysql/v4"

	"nayanjd/docket/models"
)

var Srv *server.Server

func GetSrv() *server.Server {
	if Srv == nil {
		Srv = SetupOauth()
	}

	return Srv
}

func SetupOauth() *server.Server {
	manager := manage.NewDefaultManager()

	log.Info().Msg(GetMysqlDsn())
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
		user := &models.User{}

		err = models.GetDB().First(user, "username = ?", username).Error

		if err == nil && user.ComparePassword(&password) {
			userID = *user.ID
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

	srv.SetResponseTokenHandler(
		func(w http.ResponseWriter, data map[string]interface{}, header http.Header, statusCode ...int) error {
			body := gin.H{
				"data":      data,
				"errors":    nil,
				"isSuccess": len(statusCode) > 0 && IsStatusSuccess(statusCode[0]),
				"meta":      gin.H{},
			}

			w.Header().Set("Content-Type", "application/json;charset=UTF-8")
			w.Header().Set("Cache-Control", "no-store")
			w.Header().Set("Pragma", "no-cache")

			for key := range header {
				w.Header().Set(key, header.Get(key))
			}

			status := http.StatusOK
			if len(statusCode) > 0 && statusCode[0] > 0 {
				status = statusCode[0]
			}

			w.WriteHeader(status)
			return json.NewEncoder(w).Encode(body)
		},
	)

	return srv
}
