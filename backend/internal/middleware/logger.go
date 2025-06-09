package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Logger returns a gin middleware for logging HTTP requests
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = c.GetHeader("X-Correlation-ID")
		}

		// Create a logger with request context
		log := zap.L()
		if requestID != "" {
			log = log.With(zap.String("request_id", requestID))
		}

		// Log request start
		log.Debug("Request started",
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
		)

		// Process request
		c.Next()

		// Log request completion
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		// Create log entry with response details
		fields := []zap.Field{
			zap.Int("status", statusCode),
			zap.Duration("latency", latency),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
		}

		// Add error message if any
		if errorMessage != "" {
			fields = append(fields, zap.String("error", errorMessage))
		}

		// Add user ID if available
		if userID, exists := c.Get("userID"); exists {
			fields = append(fields, zap.Uint("user_id", userID.(uint)))
		}

		// Log based on status code
		switch {
		case statusCode >= 500:
			log.Error("Request failed", fields...)
		case statusCode >= 400:
			log.Warn("Request failed", fields...)
		default:
			log.Debug("Request completed", fields...)
		}
	}
}
