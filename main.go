package main

import (
	"github.com/aprilboiz/flight-management/internal/api"
	"github.com/aprilboiz/flight-management/internal/api/handlers"
	"github.com/aprilboiz/flight-management/internal/repository"
	"github.com/aprilboiz/flight-management/internal/service"
	"github.com/aprilboiz/flight-management/pkg/database"
	"github.com/aprilboiz/flight-management/pkg/logger"
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
	logger.Init("development")
	log := logger.Get()

	log.Info("Setting up the application")
	db := database.GetDatabase()

	// Declare repositories
	flightRepo := repository.NewFlightRepository(db)
	airportRepo := repository.NewAirportRepository(db)
	planeRepo := repository.NewPlaneRepository(db)

	// Declare services
	flightService := service.NewFlightService(flightRepo, airportRepo, planeRepo)
	airportService := service.NewAirportService(airportRepo)
	planeService := service.NewPlaneService(planeRepo)

	// Declare handlers
	flightHandler := handlers.NewFlightHandler(flightService)
	airportHandler := handlers.NewAirportHandler(airportService)
	planeHandler := handlers.NewPlaneHandler(planeService)

	h := api.Handlers{
		AirportHandler: airportHandler,
		PlaneHandler:   planeHandler,
		FlightHandler:  flightHandler,
		Logger:         log,
	}
	router := gin.Default()
	api.SetupRoutes(router, h)

	log.Info("Running the server")
	err := router.Run(":8080")
	if err != nil {
		log.Error("Failed to run the server", zap.Error(err))
		return
	}
}
