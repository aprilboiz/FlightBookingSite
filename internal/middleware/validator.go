package middleware

import (
	"net/http"

	"github.com/aprilboiz/flight-management/pkg/validator"

	"github.com/gin-gonic/gin"
)

func ValidateRequest(model any) gin.HandlerFunc {
	v := validator.New()

	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(model); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if err := v.Validate(model); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validator.ValidationErrors(err)})
			c.Abort()
			return
		}

		c.Set("validatedModel", model)
		c.Next()
	}
}
