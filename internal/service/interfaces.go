package service

import (
	"github.com/aprilboiz/flight-management/internal/dto"
	"github.com/aprilboiz/flight-management/internal/models"
)

type FlightService interface {
	GetAllFlights() ([]*models.Flight, error)
	GetFlightByID(flightID string) (*models.Flight, error)
	GetFlightByCode(flightCode string) (*models.Flight, error)
	Create(flightRequest *dto.FlightRequest) (*models.Flight, error)
	Update(flightCode string, flightRequest *dto.FlightRequest) (*models.Flight, error)
	DeleteByCode(code string) error
}
