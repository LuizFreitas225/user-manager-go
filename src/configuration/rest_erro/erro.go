package manager_error

import (
	"net/http"
	"strings"
)

type RestError struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []string `json:"causes"`
}

func (e *RestError) Error() string {
	return e.Message
}

func (e *RestError) GetCauses() string {
	return strings.Join(e.Causes, ", ")
}

func NewRestError(message, err string, code int, causes []string) *RestError {
	return &RestError{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string, causes []string) *RestError {
	return &RestError{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewBadRequestValidationError(message string, causes []string) *RestError {
	return &RestError{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string, causes []string) *RestError {
	return &RestError{
		Message: message,
		Err:     "Internal Server Error",
		Code:    http.StatusInternalServerError,
		Causes:  causes,
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "Not Found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "Forbidden",
		Code:    http.StatusForbidden,
	}
}
