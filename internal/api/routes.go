package api

import (
	"github.com/aprilboiz/flight-management/internal/api/handlers"
	"github.com/aprilboiz/flight-management/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	FlightHandler handlers.FlightHandler
}

func SetupRoutes(router *gin.Engine, h Handlers) {
	// Global middleware (e.g., logging, CORS - if needed)
	// router.Use(middleware.Logger())
	// router.Use(middleware.Cors())
	router.Use(middleware.ErrorHandler())
	v1 := router.Group("/api/v1")

	flightRoutes := v1.Group("/flights")
	{
		flightRoutes.POST("", h.FlightHandler.CreateFlight)
		flightRoutes.GET("", h.FlightHandler.GetAllFlights)
		flightRoutes.GET("/:code", h.FlightHandler.GetFlightByCode)
		flightRoutes.PUT("/:code", h.FlightHandler.UpdateFlight)
		flightRoutes.DELETE("/:code", h.FlightHandler.DeleteFlightByCode)
	}
}
