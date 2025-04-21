package api

import (
	"github.com/aprilboiz/flight-management/internal/api/handlers"
	"github.com/aprilboiz/flight-management/internal/dto"
	ex "github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/internal/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handlers struct {
	FlightHandler handlers.FlightHandler
	Logger        *zap.Logger
}

func SetupRoutes(router *gin.Engine, h Handlers) {
	// Global middleware (e.g., logging, CORS - if needed)
	// router.Use(middleware.Logger())
	// router.Use(middleware.Cors())
	router.Use(middleware.ErrorHandler(h.Logger))
	router.NoRoute(func(c *gin.Context) {
		_ = c.Error(ex.NewAppError(ex.NOT_FOUND, "Not found this route!", map[string]any{
			"path":   c.Request.URL.Path,
			"method": c.Request.Method,
		}))
	})
	v1 := router.Group("/api/v1")

	flightRoutes := v1.Group("/flights")
	{
		flightRoutes.POST("", middleware.ValidateRequest(&dto.FlightRequest{}), h.FlightHandler.CreateFlight)
		flightRoutes.GET("", h.FlightHandler.GetAllFlights)
		flightRoutes.GET("/:code", h.FlightHandler.GetFlightByCode)
		flightRoutes.PUT("/:code", h.FlightHandler.UpdateFlight)
		flightRoutes.DELETE("/:code", h.FlightHandler.DeleteFlightByCode)
	}
}
