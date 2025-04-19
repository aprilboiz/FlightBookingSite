package middleware

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"

	e "github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func ErrorHandler(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add recovery to handle panics
		defer func() {
			if r := recover(); r != nil {
				// Format stack trace for better readability
				stackTrace := string(debug.Stack())
				formattedStack := formatStackTrace(stackTrace)

				if logger != nil {
					logger.Error("Recovered from panic",
						zap.Any("error", r),
						zap.String("stack", formattedStack))
				}

				// Convert panic to error response
				errMsg := fmt.Sprintf("Internal server error: %v", r)
				response := e.NewErrorResponse(
					http.StatusInternalServerError,
					http.StatusText(http.StatusInternalServerError),
					errMsg,
					nil)

				c.AbortWithStatusJSON(response.Status, response)
			}
		}()

		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			var appErr *e.AppError
			if errors.As(err, &appErr) {
				errorInfo := e.ResolveErrorType(appErr.Type)

				details := appErr.Details
				if appErr.Type == e.INTERNAL_ERROR {
					details = nil
				}
				response := e.NewErrorResponse(errorInfo.StatusCode, errorInfo.Title, appErr.Message, details)
				c.AbortWithStatusJSON(response.Status, response)
				return
			} else {
				response := e.NewErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error(), nil)
				c.AbortWithStatusJSON(response.Status, response)
				return
			}
		}
	}
}

func formatStackTrace(stack string) string {
	lines := strings.Split(stack, "\n")
	var result []string

	// Add the first line which contains "panic:" message
	if len(lines) > 0 {
		result = append(result, lines[0])
	}

	// Process stack frames in pairs
	for i := 1; i < len(lines)-1; i += 2 {
		if i+1 < len(lines) {
			// Get file/line info and function name
			fileInfo := strings.TrimSpace(lines[i+1])
			funcInfo := strings.TrimSpace(lines[i])

			// Skip runtime and standard library frames
			if strings.Contains(fileInfo, "runtime/") {
				continue
			}

			// Combine into a single line
			frame := fmt.Sprintf("%s at %s", funcInfo, fileInfo)
			result = append(result, frame)
		}
	}

	return strings.Join(result, "\n")
}
