package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
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

	PathNotFoundError = APIError{
		Code:      "PATH_NOT_FOUND",
		Http_code: http.StatusNotFound,
		Messages:  []string{"The path does not exists"},
	}

	ResourceNotFoundError = APIError{
		Code:      "RESOURCE_NOT_FOUND",
		Http_code: http.StatusNoContent,
		Messages:  []string{"The requested resource does not exists"},
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
	case error:
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Debug().Msg("Record not found")
			return &ResourceNotFoundError
		}
		log.Error().Msg("Unknown db error")
		return &InternalServerError
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
