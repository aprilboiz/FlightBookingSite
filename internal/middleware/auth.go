package middleware

import (
	"strings"

	"slices"

	"github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/pkg/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			_ = c.Error(exceptions.NewAppError(exceptions.UNAUTHORIZED, "Authorization header is required", nil))
			c.Abort()
			return
		}

		// Check if the Authorization header has the correct format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			_ = c.Error(exceptions.NewAppError(exceptions.UNAUTHORIZED, "Invalid authorization header format", nil))
			c.Abort()
			return
		}

		// Validate the token
		claims, err := auth.ValidateToken(parts[1])
		if err != nil {
			_ = c.Error(exceptions.NewAppError(exceptions.UNAUTHORIZED, "Invalid or expired token", err))
			c.Abort()
			return
		}

		// Set user information in the context
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			_ = c.Error(exceptions.NewAppError(exceptions.UNAUTHORIZED, "User role not found in context", nil))
			c.Abort()
			return
		}

		hasRole := slices.Contains(roles, userRole.(string))

		if !hasRole {
			_ = c.Error(exceptions.NewAppError(exceptions.FORBIDDEN, "Insufficient permissions", nil))
			c.Abort()
			return
		}

		c.Next()
	}
}
