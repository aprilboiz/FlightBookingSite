package api

import (
	"github.com/aprilboiz/flight-management/internal/api/handlers"
	"go.uber.org/zap"
)

type Handlers struct {
	ParameterHandler handlers.ParameterHandler
	AirportHandler   handlers.AirportHandler
	PlaneHandler     handlers.PlaneHandler
	FlightHandler    handlers.FlightHandler
	TicketHandler    handlers.TicketHandler
	UserHandler      handlers.UserHandler
	Logger           *zap.Logger
}
