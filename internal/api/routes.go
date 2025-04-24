package api

import (
	docs "github.com/aprilboiz/flight-management/docs"
	"github.com/aprilboiz/flight-management/internal/api/handlers"
	"github.com/aprilboiz/flight-management/internal/dto"
	ex "github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/internal/middleware"
	"github.com/aprilboiz/flight-management/internal/models"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"net/http"
)

type Handlers struct {
	ParameterHandler handlers.ParameterHandler
	AirportHandler   handlers.AirportHandler
	PlaneHandler     handlers.PlaneHandler
	FlightHandler    handlers.FlightHandler
	Logger           *zap.Logger
}

func SetupRoutes(router *gin.Engine, h Handlers) {
	// Global middleware (e.g., logging, CORS - if needed)
	// router.Use(middleware.Logger())
	// router.Use(middleware.Cors())
	router.Use(middleware.ErrorHandler(h.Logger))
	router.NoRoute(func(c *gin.Context) {
		_ = c.Error(&ex.AppError{
			Code:    http.StatusText(http.StatusNotFound),
			Message: "Cannot find the requested resource. Please check your request path.",
			Details: map[string]any{
				"path":   c.Request.URL.Path,
				"method": c.Request.Method,
			},
			StatusCode: http.StatusMethodNotAllowed,
		})
	})

	docs.SwaggerInfo.BasePath = "/api"
	v1 := router.Group("/api")
	{
		flightRoutes := v1.Group("/flights")
		{
			flightRoutes.POST("", middleware.ValidateRequest(&dto.FlightRequest{}), h.FlightHandler.CreateFlight)
			flightRoutes.GET("", h.FlightHandler.GetAllFlights)
			flightRoutes.GET("/:code", h.FlightHandler.GetFlightByCode)
			flightRoutes.PUT("/:code", middleware.ValidateRequest(&dto.FlightRequest{}), h.FlightHandler.UpdateFlight)
			flightRoutes.DELETE("/:code", h.FlightHandler.DeleteFlightByCode)
		}

		planeRoutes := v1.Group("/planes")
		{
			planeRoutes.GET("", h.PlaneHandler.GetAllPlanes)
			planeRoutes.GET("/:code", h.PlaneHandler.GetPlaneByCode)
		}

		airportRoutes := v1.Group("/airports")
		{
			airportRoutes.GET("", h.AirportHandler.GetAllAirports)
			airportRoutes.GET("/:code", h.AirportHandler.GetAirportByCode)
		}

		paramHandler := v1.Group("/params")
		{
			paramHandler.GET("", h.ParameterHandler.GetAllParameters)
			paramHandler.PUT("", middleware.ValidateRequest(&models.Parameter{}), h.ParameterHandler.UpdateParameters)
		}
	}

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
