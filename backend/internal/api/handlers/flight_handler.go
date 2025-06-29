package handlers

import (
	"net/http"
	"strconv"
	"time"

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
func (f *flightHandler) GetAllFlightsInList(c *gin.Context) {
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
func (f *flightHandler) GetAllFlights(c *gin.Context) {
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
func (f *flightHandler) GetFlightByCode(c *gin.Context) {
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
func (f *flightHandler) CreateFlight(c *gin.Context) {
	validatedModel, exists := c.Get("validatedModel")
	if !exists {
		_ = c.Error(e.NewAppError(e.INTERNAL, "Cannot find validated model in context", nil))
		return
	}
	flightRequest, ok := validatedModel.(*dto.FlightRequest)
	if !ok {
		_ = c.Error(e.NewAppError(e.INTERNAL, "Cannot cast validated model to FlightRequest", nil))
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
func (f *flightHandler) UpdateFlight(c *gin.Context) {
	code := c.Param("code")
	validatedModel, exists := c.Get("validatedModel")
	if !exists {
		_ = c.Error(e.NewAppError(e.INTERNAL, "Cannot find validated model in context", nil))
		return
	}
	flightRequest, ok := validatedModel.(*dto.FlightRequest)
	if !ok {
		_ = c.Error(e.NewAppError(e.INTERNAL, "Cannot cast validated model to FlightRequest", nil))
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
func (f *flightHandler) DeleteFlightByCode(c *gin.Context) {
	code := c.Param("code")
	if err := f.flightService.Delete(code); err != nil {
		_ = c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}

// GetRevenueReport godoc
//
//	@Summary		Get revenue report
//	@Description	Retrieve revenue statistics for flights in a specific month and year
//	@Tags			reports
//	@Accept			json
//	@Produce		json
//	@Param			month	query		int	true	"Month (1-12) or leave it blank for current month"
//	@Param			year	query		int	true	"Year (e.g., 2024) or leave it blank for current year"
//	@Success		200		{object}	dto.MonthlyRevenueReport
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/api/reports/revenue [get]
func (f *flightHandler) GetRevenueReport(c *gin.Context) {
	// Get year and month from query parameters
	monthStr := c.Query("month")
	yearStr := c.Query("year")

	// Parse year
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		year = time.Now().Year()
	}

	// Parse month
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		month = int(time.Now().Month())
	}

	// Get revenue report
	report, err := f.flightService.GetMonthlyRevenueReport(year, month)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, report)
}

// GetMonthlyRevenueReport godoc
//
//	@Summary		Get monthly revenue report
//	@Description	Retrieve revenue statistics for flights in a specific month
//	@Tags			reports
//	@Accept			json
//	@Produce		json
//	@Param			month	query		int	true	"Month (1-12) or leave it blank for current month"
//	@Success		200		{object}	dto.MonthlyRevenueReport
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/api/reports/revenue/monthly [get]
func (f *flightHandler) GetMonthlyRevenueReport(c *gin.Context) {
	// Get year and month from query parameters
	monthStr := c.Query("month")

	// Parse year
	year := time.Now().Year()

	// Parse month
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		month = int(time.Now().Month())
	}

	// Get revenue report
	report, err := f.flightService.GetMonthlyRevenueReport(year, month)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, report)
}

// GetYearlyRevenueReport godoc
//
//	@Summary		Get yearly revenue report
//	@Description	Retrieve revenue statistics for flights in a specific year
//	@Tags			reports
//	@Accept			json
//	@Produce		json
//	@Param			year	query		int	true	"Year (e.g., 2024) or leave it blank for current year"
//	@Success		200		{object}	dto.YearlyRevenueReport
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/api/reports/revenue/yearly [get]
func (f *flightHandler) GetYearlyRevenueReport(c *gin.Context) {
	// Get year from query parameter
	yearStr := c.Query("year")

	// Parse year
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		year = time.Now().Year()
	}

	// Get revenue report
	report, err := f.flightService.GetYearlyRevenueReport(year)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, report)
}
