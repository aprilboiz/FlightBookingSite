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

func main() {
	logger.Init("development")
	log := logger.Get()

	log.Info("Setting up the application")
	db := database.GetDatabase(log)

	flightRepo := repository.NewFlightRepository(db)
	airportRepo := repository.NewAirportRepository(db)
	ticketClassRepo := repository.NewTicketClassRepository(db)
	planeRepo := repository.NewPlaneRepository(db)
	flightService := service.NewFlightService(flightRepo, airportRepo, ticketClassRepo)
	flightHandler := handlers.NewFlightHandler(flightService)

	h := api.Handlers{FlightHandler: flightHandler}
	router := gin.Default()
	api.SetupRoutes(router, h)

	log.Info("Running the server")
	err := router.Run(":8080")
	if err != nil {
		log.Error("Failed to run the server", zap.Error(err))
		return
	}
}
