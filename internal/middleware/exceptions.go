package middleware

import (
	"errors"
	"github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			var appErr *e.AppError
			if errors.As(err, &appErr) {
				errorInfo, exists := e.ErrorType[appErr.Type]
				if !exists {
					errorInfo = e.ErrorType[e.InternalError]
				}
				details := appErr.Details
				if appErr.Type == e.InternalError {
					zap.L().Error("Internal Error", zap.String("type", appErr.Type), zap.Any("details", details))
					details = nil
				}
				response := e.NewErrorResponse(errorInfo.StatusCode, errorInfo.Title, appErr.Message, details)
				zap.L().Error(response.Error, zap.Any("details", details))
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
