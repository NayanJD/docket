package utils

import "net/http"

type APIError struct {
	Code      string   `json:"error_code"`
	Http_code int      `json:"-"`
	Messages  []string `json:"error_messages"`
}

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
