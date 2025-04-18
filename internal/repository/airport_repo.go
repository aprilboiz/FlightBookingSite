package repository

import (
	"github.com/aprilboiz/flight-management/internal/models"
	"gorm.io/gorm"
)

type airportRepository struct {
	db *gorm.DB
}

func NewAirportRepository(db *gorm.DB) AirportRepository {
	return &airportRepository{db: db}
}

func (a airportRepository) GetAll() ([]*models.Airport, error) {
	flights := make([]*models.Airport, 0)
	result := a.db.Find(&flights)
	if result.Error != nil {
		return nil, result.Error
	}
	return flights, nil
}

func (a airportRepository) GetByCode(code string) (*models.Airport, error) {
	var flight models.Airport
	result := a.db.Where("airport_code = ?", code).First(&flight)
	if result.Error != nil {
		return nil, result.Error
	}
	return &flight, nil
}

func (a airportRepository) GetByCodes(codes []string) (map[string]*models.Airport, error) {
	flights := make([]*models.Airport, 0)
	result := a.db.Where("airport_code IN ?", codes).Find(&flights)
	if result.Error != nil {
		return nil, result.Error
	}
	airportMap := make(map[string]*models.Airport)
	for _, flight := range flights {
		airportMap[flight.AirportCode] = flight
	}
	return airportMap, nil
}
