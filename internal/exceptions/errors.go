package exceptions

import (
	"fmt"
	"github.com/aprilboiz/flight-management/internal/dto"
	"net/http"
)

type ErrorInfo struct {
	StatusCode int
	Title      string
}
type AppError struct {
	Code       string
	Message    string
	Details    any
	StatusCode int
	Err        error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NotFound(entity string, identifier string) *AppError {
	return &AppError{
		Code:       NOT_FOUND,
		Message:    fmt.Sprintf("%s with identifier '%s' not found", entity, identifier),
		StatusCode: http.StatusNotFound,
	}
}

func BadRequest(message string, err error) *AppError {
	return &AppError{
		Code:       BAD_REQUEST,
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Err:        err,
	}
}

func Internal(message string, err error) *AppError {
	return &AppError{
		Code:       INTERNAL_ERROR,
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Err:        err,
	}
}

func NewErrorResponse(statusCode int, title, message string, details interface{}) *dto.ErrorResponse {
	return &dto.ErrorResponse{
		Status:  statusCode,
		Type:    title,
		Message: message,
		Details: details,
	}
}

func ResolveErrorType(errorType string) *ErrorInfo {
	errorInfo, exists := ErrorType[errorType]
	if !exists {
		errorInfo = ErrorType[INTERNAL_ERROR]
	}
	return errorInfo
}
