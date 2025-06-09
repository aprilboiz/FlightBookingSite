package api

import (
	"net/http"

	docs "github.com/aprilboiz/flight-management/docs"
	"github.com/aprilboiz/flight-management/internal/dto"
	ex "github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/internal/middleware"
	"github.com/aprilboiz/flight-management/internal/models"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine, h Handlers) {
	// Global middleware
	router.Use(middleware.Logger())
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

	docs.SwaggerInfo.BasePath = "/"
	v1 := router.Group("/api")
	{
		// Public routes
		authRoutes := v1.Group("/auth")
		{
			authRoutes.POST("/register", h.UserHandler.Register)
			authRoutes.POST("/login", h.UserHandler.Login)
		}

		// Protected routes
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// Flight routes
			flightRoutes := protected.Group("/flights")
			{
				flightRoutes.GET("", h.FlightHandler.GetAllFlights)
				flightRoutes.GET("/list", h.FlightHandler.GetAllFlightsInList)
				flightRoutes.GET("/:code", h.FlightHandler.GetFlightByCode)

				// Higher level roles
				adminFlightOps := flightRoutes.Group("")
				adminFlightOps.Use(middleware.RoleMiddleware(models.RoleAdmin, models.RoleSuperAdmin))
				{
					adminFlightOps.POST("", middleware.ValidateRequest(&dto.FlightRequest{}), h.FlightHandler.CreateFlight)
					adminFlightOps.PUT("/:code", middleware.ValidateRequest(&dto.FlightRequest{}), h.FlightHandler.UpdateFlight)
					adminFlightOps.DELETE("/:code", h.FlightHandler.DeleteFlightByCode)
				}
			}

			// Plane routes
			planeRoutes := protected.Group("/planes")
			{
				planeRoutes.GET("", h.PlaneHandler.GetAllPlanes)
				planeRoutes.GET("/:code", h.PlaneHandler.GetPlaneByCode)
			}

			// Airport routes
			airportRoutes := protected.Group("/airports")
			{
				airportRoutes.GET("", h.AirportHandler.GetAllAirports)
				airportRoutes.GET("/:code", h.AirportHandler.GetAirportByCode)
			}

			// Parameter routes
			adminRoutes := protected.Group("")
			adminRoutes.Use(middleware.RoleMiddleware(models.RoleSuperAdmin))
			{
				paramHandler := adminRoutes.Group("/params")
				{
					paramHandler.GET("", h.ParameterHandler.GetAllParameters)
					paramHandler.PUT("", middleware.ValidateRequest(&models.Parameter{}), h.ParameterHandler.UpdateParameters)
				}
			}

			// Report routes
			reportRoutes := protected.Group("")
			reportRoutes.Use(middleware.RoleMiddleware(models.RoleAdmin, models.RoleSuperAdmin))
			{
				reportHandler := reportRoutes.Group("/reports")
				{
					reportHandler.GET("/revenue/monthly", h.FlightHandler.GetMonthlyRevenueReport)
					reportHandler.GET("/revenue/yearly", h.FlightHandler.GetYearlyRevenueReport)
					reportHandler.GET("/revenue", h.FlightHandler.GetRevenueReport)
				}
			}

			// Ticket operations
			ticketRoutes := protected.Group("/tickets")
			{
				ticketRoutes.GET("", h.TicketHandler.GetAllTickets)
				ticketRoutes.GET("/:id", h.TicketHandler.GetTicketByID)
				ticketRoutes.POST("", middleware.ValidateRequest(&dto.TicketRequest{}), h.TicketHandler.CreateTicket)
				ticketRoutes.PUT("/:id/status", middleware.ValidateRequest(&dto.TicketStatusUpdateRequest{}), h.TicketHandler.UpdateTicketStatus)
				ticketRoutes.DELETE("/:id", h.TicketHandler.DeleteTicket)
				ticketRoutes.GET("/statuses", h.TicketHandler.GetTicketStatuses)
				ticketRoutes.GET("/booking-types", h.TicketHandler.GetBookingTypes)
			}
		}
	}

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
