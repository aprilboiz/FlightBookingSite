package handlers

import (
	"github.com/aprilboiz/flight-management/internal/models"
	"github.com/aprilboiz/flight-management/internal/service"
	"github.com/gin-gonic/gin"
)

type flightHandler struct {
	flightService service.FlightService
}

func NewFlightHandler(flightService service.FlightService) FlightHandler {
	return &flightHandler{flightService: flightService}
}

func (f flightHandler) CreateFlight(c *gin.Context) (*models.Flight, error) {
	//TODO implement me
	panic("implement me")
}

func (f flightHandler) GetAllFlights(c *gin.Context) ([]*models.Flight, error) {
	//TODO implement me
	panic("implement me")
}

func (f flightHandler) GetFlightByCode(c *gin.Context) (*models.Flight, error) {
	//TODO implement me
	panic("implement me")
}

func (f flightHandler) UpdateFlight(c *gin.Context) (*models.Flight, error) {
	//TODO implement me
	panic("implement me")
}

func (f flightHandler) DeleteFlight(c *gin.Context) error {
	//TODO implement me
	panic("implement me")
}
