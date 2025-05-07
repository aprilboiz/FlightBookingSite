package handlers

import (
	"net/http"

	"github.com/aprilboiz/flight-management/internal/dto"
	e "github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/internal/service"
	"github.com/gin-gonic/gin"
)

func NewFlightHandler(flightService service.FlightService) FlightHandler {
	return &flightHandler{flightService: flightService}
}

type flightHandler struct {
	flightService service.FlightService
}

// GetAllFlightsInList godoc
//
//	@Summary		Get all flights in list format
//	@Description	Retrieve a list of all flights with simplified information
//	@Tags			flights
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		dto.FlightListResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/api/flights/list [get]
func (f flightHandler) GetAllFlightsInList(c *gin.Context) {
	flights, err := f.flightService.GetAllFlightsInList()
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, flights)
}

// GetAllFlights godoc
//
//	@Summary		Get all flights
//	@Description	Retrieve a list of all flights
//	@Tags			flights
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		dto.FlightResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/api/flights [get]
func (f flightHandler) GetAllFlights(c *gin.Context) {
	flights, err := f.flightService.GetAllFlights()
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, flights)
}

// GetFlightByCode godoc
//
//	@Summary		Get flight by code
//	@Description	Retrieve a flight by its unique code
//	@Tags			flights
//	@Accept			json
//	@Produce		json
//	@Param			code	path		string	true	"Flight Code"
//	@Success		200		{object}	dto.FlightResponseDetailed
//	@Failure		404		{object}	dto.ErrorResponse
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/api/flights/{code} [get]
func (f flightHandler) GetFlightByCode(c *gin.Context) {
	code := c.Param("code")
	flight, err := f.flightService.GetFlightByCode(code)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, flight)
}

// CreateFlight godoc
//
//	@Summary		Create a new flight
//	@Description	Create a new flight with the provided information
//	@Tags			flights
//	@Accept			json
//	@Produce		json
//	@Param			flight	body		dto.FlightRequest	true	"Flight information"
//	@Success		201		{object}	dto.FlightResponse
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/api/flights [post]
func (f flightHandler) CreateFlight(c *gin.Context) {
	validatedModel, exists := c.Get("validatedModel")
	if !exists {
		_ = c.Error(e.NewAppError(e.INTERNAL_ERROR, "Cannot find validated model in context", nil))
		return
	}
	flightRequest, ok := validatedModel.(*dto.FlightRequest)
	if !ok {
		_ = c.Error(e.NewAppError(e.INTERNAL_ERROR, "Cannot cast validated model to FlightRequest", nil))
		return
	}
	flightResponse, err := f.flightService.Create(flightRequest)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, flightResponse)
}

// UpdateFlight godoc
//
//	@Summary		Update a flight
//	@Description	Update an existing flight with the provided information
//	@Tags			flights
//	@Accept			json
//	@Produce		json
//	@Param			code	path		string				true	"Flight Code"
//	@Param			flight	body		dto.FlightRequest	true	"Flight information"
//	@Success		200		{object}	dto.FlightResponse
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		404		{object}	dto.ErrorResponse
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/api/flights/{code} [put]
func (f flightHandler) UpdateFlight(c *gin.Context) {
	code := c.Param("code")
	validatedModel, exists := c.Get("validatedModel")
	if !exists {
		_ = c.Error(e.NewAppError(e.INTERNAL_ERROR, "Cannot find validated model in context", nil))
		return
	}
	flightRequest, ok := validatedModel.(*dto.FlightRequest)
	if !ok {
		_ = c.Error(e.NewAppError(e.INTERNAL_ERROR, "Cannot cast validated model to FlightRequest", nil))
		return
	}

	flightResponse, err := f.flightService.Update(code, flightRequest)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, flightResponse)
}

// DeleteFlightByCode godoc
//
//	@Summary		Delete a flight
//	@Description	Delete a flight by its unique code
//	@Tags			flights
//	@Accept			json
//	@Produce		json
//	@Param			code	path	string	true	"Flight Code"
//	@Success		204		"No Content"
//	@Failure		404		{object}	dto.ErrorResponse
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/api/flights/{code} [delete]
func (f flightHandler) DeleteFlightByCode(c *gin.Context) {
	code := c.Param("code")
	if err := f.flightService.Delete(code); err != nil {
		_ = c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}
