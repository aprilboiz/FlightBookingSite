package middleware

import (
	"net/http"

	"github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/pkg/validator"

	"github.com/gin-gonic/gin"
)

func ValidateRequest(model any) gin.HandlerFunc {
	v := validator.New()

	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(model); err != nil {
			response := exceptions.NewErrorResponse(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				"Invalid request body",
				err.Error(),
			)
			c.JSON(http.StatusBadRequest, response)
			c.Abort()
			return
		}

		if err := v.Validate(model); err != nil {
			validationErrors := validator.ValidationErrors(err)
			response := exceptions.NewErrorResponse(
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				"Validation failed",
				validationErrors,
			)
			c.JSON(http.StatusBadRequest, response)
			c.Abort()
			return
		}

		c.Set("validatedModel", model)
		c.Next()
	}
}
