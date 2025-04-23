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

type PlaneHandler interface {
	GetAllPlanes(c *gin.Context)
	GetPlaneByCode(c *gin.Context)
}

type AirportHandler interface {
	GetAllAirports(c *gin.Context)
	GetAirportByCode(c *gin.Context)
}
