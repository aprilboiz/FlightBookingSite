package service

import (
	"github.com/aprilboiz/flight-management/internal/dto"
	"github.com/aprilboiz/flight-management/internal/models"
)

type FlightService interface {
	GetAllFlights() ([]*dto.FlightResponse, error)
	GetFlightByID(flightID string) (*dto.FlightResponse, error)
	GetFlightByCode(flightCode string) (*dto.FlightResponse, error)
	Create(flightRequest *dto.FlightRequest) (*dto.FlightResponse, error)
	Update(flightCode string, flightRequest *dto.FlightRequest) (*dto.FlightResponse, error)
	DeleteByCode(code string) error
}

type AirportService interface {
	GetAllAirports() ([]*models.Airport, error)
	GetAirportByCode(code string) (*models.Airport, error)
	GetAirportsByCodes(codes []string) (map[string]*models.Airport, error)
}

type FlightCodeGenerator interface {
	Generate() (string, error)
}
