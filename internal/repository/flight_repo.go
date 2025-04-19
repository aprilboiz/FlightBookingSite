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
	var flights []*models.Flight
	result := f.db.
		Preload("DepartureAirport").
		Preload("ArrivalAirport").
		Preload("Plane").
		Preload("IntermediateStops.Airport").
		Find(&flights)
	if result.Error != nil {
		return nil, result.Error
	}
	return flights, nil
}

func (f fightRepository) GetByID(id int) (*models.Flight, error) {
	var flight models.Flight
	result := f.db.
		Preload("DepartureAirport").
		Preload("ArrivalAirport").
		Preload("Plane").
		Preload("IntermediateStops.Airport").
		Where("id = ?", id).
		First(&flight)
	if result.Error != nil {
		return nil, result.Error
	}

	return &flight, nil
}

func (f fightRepository) GetByCode(code string) (*models.Flight, error) {
	var flight models.Flight
	result := f.db.
		Preload("DepartureAirport").
		Preload("ArrivalAirport").
		Preload("Plane").
		Preload("IntermediateStops.Airport").
		Where("flight_code = ?", code).
		First(&flight)
	if result.Error != nil {
		return nil, result.Error
	}
	return &flight, nil
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

func (f fightRepository) CreateIntermediateStops(stops []*models.IntermediateStop) ([]*models.IntermediateStop, error) {
	result := f.db.Create(&stops)
	if result.Error != nil {
		return nil, result.Error
	}
	return stops, nil
}
