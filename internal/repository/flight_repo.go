package repository

import (
	"errors"
	"github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/internal/models"
	"gorm.io/gorm"
	"strconv"
)

type flightRepository struct {
	db *gorm.DB
}

func NewFlightRepository(db *gorm.DB) FlightRepository {
	return &flightRepository{db: db}
}

func (f flightRepository) GetAll() ([]*models.Flight, error) {
	var flights []*models.Flight
	result := f.db.
		Preload("DepartureAirport").
		Preload("ArrivalAirport").
		Preload("Plane").
		Preload("IntermediateStops.Airport").
		Find(&flights)
	if result.Error != nil {
		return nil, exceptions.Internal("failed to get all flights", result.Error)
	}
	return flights, nil
}

func (f flightRepository) GetByID(id uint) (*models.Flight, error) {
	var flight models.Flight
	result := f.db.
		Preload("DepartureAirport").
		Preload("ArrivalAirport").
		Preload("Plane").
		Preload("IntermediateStops.Airport").
		Where("id = ?", id).
		First(&flight)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, exceptions.NotFound("flight", strconv.Itoa(int(id)))
		}
		return nil, exceptions.Internal("failed to get flight by id", result.Error)
	}

	return &flight, nil
}

func (f flightRepository) GetByCode(code string) (*models.Flight, error) {
	var flight models.Flight
	result := f.db.
		Preload("DepartureAirport").
		Preload("ArrivalAirport").
		Preload("Plane").
		Preload("IntermediateStops.Airport").
		Where("flight_code = ?", code).
		First(&flight)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, exceptions.NotFound("flight", code)
		}
		return nil, exceptions.Internal("failed to get flight by code", result.Error)
	}
	return &flight, nil
}

func (f flightRepository) Create(flight *models.Flight) (*models.Flight, error) {
	result := f.db.Create(flight)
	if result.Error != nil {
		return nil, exceptions.Internal("failed to create flight", result.Error)
	}
	return flight, nil
}

func (f flightRepository) Update(flight *models.Flight) (*models.Flight, error) {
	//TODO implement me
	panic("implement me")
}

func (f flightRepository) Delete(flight *models.Flight) error {
	//TODO implement me
	panic("implement me")
}

func (f flightRepository) CreateIntermediateStops(stops []*models.IntermediateStop) ([]*models.IntermediateStop, error) {
	result := f.db.Create(&stops)
	if result.Error != nil {
		return nil, exceptions.Internal("failed to create intermediate stops", result.Error)
	}
	return stops, nil
}

func (f flightRepository) GetAvailableSeats(flightCode string) ([]*models.Seat, error) {
	var seats []*models.Seat
	result := f.db.
		Preload("TicketClass").
		Preload("Tickets").
		Joins("JOIN flights ON flights.plane_id = seats.plane_id").
		Where("flights.flight_code = ?", flightCode).
		Find(&seats)
	if result.Error != nil {
		return nil, exceptions.Internal("failed to get available seats", result.Error)
	}
	return seats, nil
}
