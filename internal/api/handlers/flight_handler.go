package handlers

import (
	"github.com/aprilboiz/flight-management/internal/dto"
	e "github.com/aprilboiz/flight-management/internal/exceptions"
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
	flights, err := f.flightService.GetAllFlights()
	if err != nil {
		_ = c.Error(e.NewAppError(e.BAD_REQUEST, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, flights)
}

func (f flightHandler) GetFlightByCode(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (f flightHandler) CreateFlight(c *gin.Context) {
	validatedModel, exists := c.Get("validatedModel")
	if !exists {
		_ = c.Error(e.NewAppError(e.BAD_REQUEST, "validated model not found", nil))
		return
	}
	flightRequest, ok := validatedModel.(*dto.FlightRequest)
	if !ok {
		_ = c.Error(e.NewAppError(e.BAD_REQUEST, "invalid flight request body", nil))
		return
	}
	flightResponse, err := f.flightService.Create(flightRequest)
	if err != nil {
		_ = c.Error(e.NewAppError(e.BAD_REQUEST, err.Error(), nil))
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
