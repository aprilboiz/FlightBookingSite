package main

import (
	"fmt"
	"github.com/aprilboiz/flight-management/pkg/config"
	"os"
	"os/signal"
	"syscall"
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

// @title Flight Management API
// @version 1.0
// @description API for flight management system
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.ruaairline.com/support
// @contact.email support@ruaairline.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
// @schemes http https

func main() {
	// Initialize logger
	logger.Init(config.GetConfig().Environment)
	log := logger.Get()

	// Set up clean shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("Shutting down...")
		err := logger.Sync()
		if err != nil {
			return
		}
		os.Exit(0)
	}()

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

	h := api.Handlers{
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
	api.SetupRoutes(router, h)

	// Start server
	log.Info("Running the server on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Error("Failed to run the server", zap.Error(err))
	}
}
