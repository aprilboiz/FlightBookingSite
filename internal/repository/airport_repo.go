package repository

import (
	"errors"
	"github.com/aprilboiz/flight-management/internal/exceptions"
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
	airports := make([]*models.Airport, 0)

	if err := a.db.Find(&airports).Error; err != nil {
		return nil, exceptions.Internal("failed to get all airports", err)
	}

	return airports, nil
}

func (a airportRepository) GetByCode(code string) (*models.Airport, error) {
	var flight models.Airport
	result := a.db.Where("airport_code = ?", code).First(&flight)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, exceptions.NotFound("airport", code)
		}
		return nil, exceptions.Internal("failed to get airport by code", result.Error)
	}
	return &flight, nil
}

func (a airportRepository) GetByCodes(codes []string) (map[string]*models.Airport, error) {
	airports := make([]*models.Airport, 0, len(codes))
	result := a.db.Where("airport_code IN ?", codes).Find(&airports)

	if result.Error != nil {
		return nil, exceptions.Internal("failed to get airports by codes", result.Error)
	}

	airportMap := make(map[string]*models.Airport, len(airports))
	for _, airport := range airports {
		airportMap[airport.AirportCode] = airport
	}

	return airportMap, nil
}
