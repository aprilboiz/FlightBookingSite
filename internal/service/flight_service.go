package service

import (
	"github.com/aprilboiz/flight-management/internal/dto"
	"github.com/aprilboiz/flight-management/internal/models"
	"github.com/aprilboiz/flight-management/internal/repository"
)

type flightService struct {
	flightRepo repository.FlightRepository
}

func NewFlightService(flightRepo repository.FlightRepository) FlightService {
	return &flightService{flightRepo: flightRepo}
}

func (f flightService) GetAllFlights() ([]*models.Flight, error) {
	//TODO implement me
	panic("implement me")
}

func (f flightService) GetFlightByID(flightID string) (*models.Flight, error) {
	//TODO implement me
	panic("implement me")
}

func (f flightService) GetFlightByCode(flightCode string) (*models.Flight, error) {
	//TODO implement me
	panic("implement me")
}

func (f flightService) Create(flightRequest *dto.FlightRequest) (*models.Flight, error) {
	//TODO implement me
	panic("implement me")
}

func (f flightService) Update(flightCode string, flightRequest *dto.FlightRequest) (*models.Flight, error) {
	//TODO implement me
	panic("implement me")
}

func (f flightService) DeleteByCode(code string) error {
	//TODO implement me
	panic("implement me")
}
