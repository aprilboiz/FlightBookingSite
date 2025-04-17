package handlers

import (
	"github.com/aprilboiz/flight-management/internal/service"
	"github.com/gin-gonic/gin"
)

type flightHandler struct {
	flightService service.FlightService
}

func NewFlightHandler(flightService service.FlightService) FlightHandler {
	return &flightHandler{flightService: flightService}
}

func (f flightHandler) GetAllFlights(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (f flightHandler) GetFlightByCode(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (f flightHandler) CreateFlight(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (f flightHandler) UpdateFlight(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (f flightHandler) DeleteFlightByCode(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
