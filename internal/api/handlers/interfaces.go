package handlers

import (
	"github.com/aprilboiz/flight-management/internal/models"
	"github.com/gin-gonic/gin"
)

type FlightHandler interface {
	GetAllFlights(c *gin.Context) ([]*models.Flight, error)
	GetFlightByCode(c *gin.Context) (*models.Flight, error)
	CreateFlight(c *gin.Context) (*models.Flight, error)
	UpdateFlight(c *gin.Context) (*models.Flight, error)
	DeleteFlight(c *gin.Context) error
}
