package service

import (
	"github.com/aprilboiz/flight-management/internal/models"
	"github.com/aprilboiz/flight-management/internal/repository"
)

type airportService struct {
	airportRepo repository.AirportRepository
}

func NewAirportService(airportRepo repository.AirportRepository) AirportService {
	return &airportService{airportRepo: airportRepo}
}

func (a airportService) GetAllAirports() ([]*models.Airport, error) {
	airports, err := a.airportRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return airports, nil
}

func (a airportService) GetAirportByCode(code string) (*models.Airport, error) {
	airport, err := a.airportRepo.GetByCode(code)
	if err != nil {
		return nil, err
	}
	return airport, nil
}

func (a airportService) GetAirportsByCodes(codes []string) (map[string]*models.Airport, error) {
	airports, err := a.airportRepo.GetByCodes(codes)
	if err != nil {
		return nil, err
	}
	return airports, nil
}
