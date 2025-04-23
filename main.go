package main

import (
	"time"

	"github.com/aprilboiz/flight-management/internal/api"
	"github.com/aprilboiz/flight-management/internal/api/handlers"
	"github.com/aprilboiz/flight-management/internal/repository"
	"github.com/aprilboiz/flight-management/internal/service"
	"github.com/aprilboiz/flight-management/pkg/database"
	"github.com/aprilboiz/flight-management/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	logger.Init("development")
	log := logger.Get()

	log.Info("Setting up the application")

	// Initialize database connection
	db := database.GetDatabase()

	// Repositories
	flightRepo := repository.NewFlightRepository(db)
	airportRepo := repository.NewAirportRepository(db)
	planeRepo := repository.NewPlaneRepository(db)

	// Services
	flightService := service.NewFlightService(flightRepo, airportRepo, planeRepo)
	airportService := service.NewAirportService(airportRepo)
	planeService := service.NewPlaneService(planeRepo)

	// Handlers
	flightHandler := handlers.NewFlightHandler(flightService)
	airportHandler := handlers.NewAirportHandler(airportService)
	planeHandler := handlers.NewPlaneHandler(planeService)

	handlers := api.Handlers{
		AirportHandler: airportHandler,
		PlaneHandler:   planeHandler,
		FlightHandler:  flightHandler,
		Logger:         log,
	}

	// Create Gin router
	router := gin.Default()

	// ✅ Enable CORS for frontend (e.g. React at localhost:5173)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup routes
	api.SetupRoutes(router, handlers)

	// Start server
	log.Info("Running the server on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Error("Failed to run the server", zap.Error(err))
	}
}
