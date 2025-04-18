package handlers

import (
	"github.com/aprilboiz/flight-management/internal/dto"
	"github.com/aprilboiz/flight-management/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
	validatedModel, exists := c.Get("validatedModel")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validated model not found"})
		return
	}
	flightRequest, ok := validatedModel.(*dto.FlightRequest)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid validated model"})
		return
	}
	flightResponse, err := f.flightService.Create(flightRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, flightResponse)
}

func (f flightHandler) UpdateFlight(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (f flightHandler) DeleteFlightByCode(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
