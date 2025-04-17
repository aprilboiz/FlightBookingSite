package handlers

import (
	"github.com/gin-gonic/gin"
)

type FlightHandler interface {
	GetAllFlights(c *gin.Context)
	GetFlightByCode(c *gin.Context)
	CreateFlight(c *gin.Context)
	UpdateFlight(c *gin.Context)
	DeleteFlightByCode(c *gin.Context)
}
