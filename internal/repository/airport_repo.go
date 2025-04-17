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

func (a airportRepository) GetByCode(code string) (*models.Airport, error) {
	//TODO implement me
	panic("implement me")
}

func (a airportRepository) GetByCodes(codes []string) (map[string]*models.Airport, error) {
	//TODO implement me
	panic("implement me")
}
