package exceptions

import (
	"fmt"
	"github.com/aprilboiz/flight-management/internal/dto"
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

func NewAppError(errType, message string, details interface{}) *AppError {
	errInfo := ResolveErrorType(errType)
	return &AppError{
		Code:       errInfo.Title,
		Message:    message,
		Details:    details,
		StatusCode: errInfo.StatusCode,
	}
}

func NotFound(entity string, identifier string) *AppError {
	errInfo := ResolveErrorType(NOT_FOUND)
	return &AppError{
		Code:       errInfo.Title,
		Message:    fmt.Sprintf("%s with identifier '%s' not found", entity, identifier),
		StatusCode: errInfo.StatusCode,
	}
}

func BadRequest(message string, err error) *AppError {
	errInfo := ResolveErrorType(BAD_REQUEST)
	return &AppError{
		Code:       errInfo.Title,
		Message:    message,
		StatusCode: errInfo.StatusCode,
		Err:        err,
	}
}

func Internal(message string, err error) *AppError {
	errInfo := ResolveErrorType(INTERNAL_ERROR)
	return &AppError{
		Code:       errInfo.Title,
		Message:    message,
		StatusCode: errInfo.StatusCode,
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
