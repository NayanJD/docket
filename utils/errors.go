package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

type APIError struct {
	Code      string   `json:"error_code"`
	Http_code int      `json:"-"`
	Messages  []string `json:"error_messages"`
}

const (
	ErrorTypeDB gin.ErrorType = 2
)

var (
	InternalServerError = APIError{
		Code:      "INTERNAL_ERROR",
		Http_code: http.StatusInternalServerError,
		Messages:  []string{"Something went wrong"},
	}
)

func CreateUnprocessableEntityError(errors []string) *APIError {
	return &APIError{
		Code:      "UNPROCESSABLE_ENTITY",
		Http_code: http.StatusUnprocessableEntity,
		Messages:  errors,
	}
}

func CreateConflictError(errors []string) *APIError {
	return &APIError{
		Code:      "CONFLICT",
		Http_code: http.StatusConflict,
		Messages:  errors,
	}
}

func CreateDbError(err error) *APIError {
	switch v := err.(type) {
	case *mysql.MySQLError:
		log.Debug().Msg(fmt.Sprintf("Error code: %v", v.Number))
		return HandleMysqlError(v)
	default:
		log.Error().Msg("Unknown db error")
		return &InternalServerError
	}
}

func HandleMysqlError(err *mysql.MySQLError) *APIError {
	switch err.Number {
	case 1062:
		return CreateConflictError([]string{err.Message})
	default:
		return &InternalServerError
	}
}
