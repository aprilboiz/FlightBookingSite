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

	tests := v1.Group("/test")
	{
		tests.GET("/", handlers.TestHandler)
		tests.POST("/", handlers.TestPostHandler)
	}

	flightRoutes := v1.Group("/flights")
	{
		flightRoutes.POST("", h.FlightHandler.CreateFlight)               // POST /api/flights
		flightRoutes.GET("", h.FlightHandler.GetAllFlights)               // GET /api/flights
		flightRoutes.GET("/:code", h.FlightHandler.GetFlightByCode)       // GET /api/flights/{code}
		flightRoutes.PUT("/:code", h.FlightHandler.UpdateFlight)          // PUT /api/flights/{code}
		flightRoutes.DELETE("/:code", h.FlightHandler.DeleteFlightByCode) // DELETE /api/flights/{code}
		// Removed GetFlightByID as GetFlightByCode seems to be the primary identifier used in service
	}
}
