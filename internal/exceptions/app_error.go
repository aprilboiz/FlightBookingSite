package exceptions

import "github.com/aprilboiz/flight-management/internal/dto"

type ErrorInfo struct {
	StatusCode int
	Title      string
}
type AppError struct {
	Type    string
	Message string
	Details interface{}
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(errorType, message string, details interface{}) *AppError {
	return &AppError{
		Type:    errorType,
		Message: message,
		Details: details,
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
