package repository

import (
	"github.com/aprilboiz/flight-management/internal/models"
	"gorm.io/gorm"
)

type fightRepository struct {
	db *gorm.DB
}

func NewFlightRepository(db *gorm.DB) FlightRepository {
	return &fightRepository{db: db}
}

func (f fightRepository) GetAll() ([]*models.Flight, error) {
	//TODO implement me
	panic("implement me")
}

func (f fightRepository) GetByID(id int) (*models.Flight, error) {
	var flight models.Flight
	result := f.db.First(&flight, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &flight, nil
}

func (f fightRepository) GetByCode(code string) (*models.Flight, error) {
	//TODO implement me
	panic("implement me")
}

func (f fightRepository) Create(flight *models.Flight) (*models.Flight, error) {
	result := f.db.Create(flight)
	if result.Error != nil {
		return nil, result.Error
	}
	return flight, nil
}

func (f fightRepository) Update(flight *models.Flight) (*models.Flight, error) {
	//TODO implement me
	panic("implement me")
}

func (f fightRepository) Delete(flight *models.Flight) error {
	//TODO implement me
	panic("implement me")
}
