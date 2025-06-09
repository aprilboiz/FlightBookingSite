package handlers

import (
	"github.com/aprilboiz/flight-management/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewAirportHandler(airportService service.AirportService) AirportHandler {
	if airportService == nil {
		panic("Missing required airport service")
	}
	return &airportHandler{airportService: airportService}
}

type airportHandler struct {
	airportService service.AirportService
}

// GetAllAirports godoc
//	@Summary		Get all airports
//	@Description	Retrieve a list of all airports
//	@Tags			airports
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		dto.AirportResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/api/airports [get]
func (h *airportHandler) GetAllAirports(c *gin.Context) {
	airports, err := h.airportService.GetAllAirports()
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, airports)
}

// GetAirportByCode godoc
//	@Summary		Get airport by code
//	@Description	Retrieve an airport by its unique code
//	@Tags			airports
//	@Accept			json
//	@Produce		json
//	@Param			code	path		string	true	"Airport Code"
//	@Success		200		{object}	dto.AirportResponse
//	@Failure		404		{object}	dto.ErrorResponse
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/api/airports/{code} [get]
func (h *airportHandler) GetAirportByCode(c *gin.Context) {
	code := c.Param("code")
	airport, err := h.airportService.GetAirportByCode(code)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, airport)
}
