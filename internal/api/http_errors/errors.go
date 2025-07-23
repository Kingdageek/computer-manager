package http_errors

import (
	"net/http"
)

type ErrorType string

const (
	ErrorTypeInternal   ErrorType = "INTERNAL"
	ErrorTypeValidation ErrorType = "VALIDATION"
	ErrorTypeNotFound   ErrorType = "NOT_FOUND"
	ErrorTypeBadRequest ErrorType = "BAD_REQUEST"
)

type HttpError struct {
	Type     ErrorType           `json:"-"`
	Code     string              `json:"code"`
	Message  string              `json:"message"`
	HTTPCode int                 `json:"-"`
	Detail   string              `json:"detail,omitempty"`
	Raw      error               `json:"-"`
	Errors   map[string][]string `json:"errors,omitempty"`
}

func (e *HttpError) Error() string {
	return e.Message
}

// NewError creates a new HttpError
func NewError(errType ErrorType, code string, message string) *HttpError {
	return &HttpError{
		Type:     errType,
		Code:     code,
		Message:  message,
		HTTPCode: getHTTPCode(errType),
	}
}

func NewNotFoundError(message string) *HttpError {
	return &HttpError{
		Type:     ErrorTypeNotFound,
		Code:     "NOT_FOUND",
		Message:  message,
		HTTPCode: http.StatusNotFound,
	}
}

func NewValidationError(errors map[string][]string) *HttpError {
	errorMessage := "The given data was invalid"
	for k := range errors {
		errorMessage = errors[k][0]
		break
	}

	return &HttpError{
		Type:     ErrorTypeValidation,
		Message:  errorMessage,
		HTTPCode: http.StatusUnprocessableEntity,
		Errors:   errors,
	}
}

func NewInternalError(err error) *HttpError {
	return &HttpError{
		Type:     ErrorTypeInternal,
		Code:     "INTERNAL_ERROR",
		Message:  "An internal error occurred",
		Detail:   err.Error(),
		HTTPCode: http.StatusInternalServerError,
		Raw:      err,
	}
}

func NewBadRequestError(message string) *HttpError {
	return &HttpError{
		Type:     ErrorTypeBadRequest,
		Code:     "BAD_REQUEST",
		Message:  message,
		HTTPCode: http.StatusBadRequest,
	}
}

// getHTTPCode maps error types to HTTP status codes
func getHTTPCode(errType ErrorType) int {
	switch errType {
	case ErrorTypeInternal:
		return http.StatusInternalServerError
	case ErrorTypeValidation:
		return http.StatusBadRequest
	case ErrorTypeNotFound:
		return http.StatusNotFound
	case ErrorTypeBadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
