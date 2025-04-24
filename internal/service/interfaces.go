package service

import (
	"github.com/aprilboiz/flight-management/internal/dto"
	"github.com/aprilboiz/flight-management/internal/models"
)

type FlightService interface {
	GetAllFlights() ([]*dto.FlightResponse, error)
	GetFlightByID(flightID int) (*dto.FlightResponse, error)
	GetFlightByCode(flightCode string) (*dto.FlightResponse, error)
	Create(flightRequest *dto.FlightRequest) (*dto.FlightResponse, error)
	Update(flightCode string, flightRequest *dto.FlightRequest) (*dto.FlightResponse, error)
	DeleteByCode(code string) error
}

type AirportService interface {
	GetAllAirports() ([]*dto.AirportResponse, error)
	GetAirportByCode(code string) (*dto.AirportResponse, error)
	GetAirportsByCodes(codes []string) (map[string]*dto.AirportResponse, error)
}

type PlaneService interface {
	GetAllPlanes() ([]*dto.PlaneResponse, error)
	GetPlaneByCode(code string) (*dto.PlaneResponseDetails, error)
}

type ParameterService interface {
	GetAllParams() (*models.Parameter, error)
	UpdateParams(params *models.Parameter) (*models.Parameter, error)
}

type FlightCodeGenerator interface {
	Generate() (string, error)
}
