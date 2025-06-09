package main

import (
	"time"

	"github.com/aprilboiz/flight-management/pkg/config"
	"github.com/aprilboiz/flight-management/pkg/logger"

	"github.com/aprilboiz/flight-management/internal/api"
	"github.com/aprilboiz/flight-management/internal/api/handlers"
	"github.com/aprilboiz/flight-management/internal/repository"
	"github.com/aprilboiz/flight-management/internal/service"
	"github.com/aprilboiz/flight-management/pkg/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//	@title			Flight Management API
//	@version		1.0
//	@description	API for flight management system
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.ruaairline.com/support
//	@contact.email	support@ruaairline.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api
//	@schemes	http https

func main() {
	// Initialize logger
	log := logger.InitLogger(config.GetConfig().Environment)
	defer func(log *zap.Logger) {
		err := log.Sync()
		if err != nil {
			log.Error("Failed to sync logger", zap.Error(err))
		}
	}(log)

	log.Info("Starting application")

	// Initialize database connection
	db := database.GetDatabase()

	// Repositories
	paramRepo := repository.NewParameterRepository(db)
	flightRepo := repository.NewFlightRepository(db)
	airportRepo := repository.NewAirportRepository(db)
	planeRepo := repository.NewPlaneRepository(db)
	ticketRepo := repository.NewTicketRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Services
	paramService := service.NewParamService(paramRepo)
	flightService := service.NewFlightService(flightRepo, airportRepo, planeRepo, paramRepo, ticketRepo)
	airportService := service.NewAirportService(airportRepo)
	planeService := service.NewPlaneService(planeRepo)
	ticketService := service.NewTicketService(ticketRepo, flightRepo, planeRepo, paramRepo)
	userService := service.NewUserService(userRepo)

	// Initialize scheduler service
	schedulerService := service.NewSchedulerService(ticketRepo, flightRepo)

	// Start the place order cancellation job
	schedulerService.StartPlaceOrderCancellationJob()

	// Handlers
	paramHandler := handlers.NewParameterHandler(paramService)
	flightHandler := handlers.NewFlightHandler(flightService)
	airportHandler := handlers.NewAirportHandler(airportService)
	planeHandler := handlers.NewPlaneHandler(planeService)
	ticketHandler := handlers.NewTicketHandler(ticketService)
	userHandler := handlers.NewUserHandler(userService, log)

	h := api.Handlers{
		ParameterHandler: paramHandler,
		AirportHandler:   airportHandler,
		PlaneHandler:     planeHandler,
		FlightHandler:    flightHandler,
		TicketHandler:    ticketHandler,
		UserHandler:      userHandler,
		Logger:           log,
	}

	// Create Gin router
	router := gin.Default()

	// ✅ Enable CORS for frontend (e.g. React at localhost:5173)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup other routes
	api.SetupRoutes(router, h)

	// Start server
	log.Info("Running the server on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Error("Failed to run the server", zap.Error(err))
	}
}
