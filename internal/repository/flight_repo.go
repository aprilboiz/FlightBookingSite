package repository

import (
	"errors"
	"strconv"
	"time"

	"github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/internal/models"
	"gorm.io/gorm"
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
	result := f.db.Save(flight)
	if result.Error != nil {
		return nil, exceptions.Internal("failed to update flight", result.Error)
	}
	return flight, nil
}

func (f flightRepository) Delete(flight *models.Flight) error {
	result := f.db.Delete(flight)
	if result.Error != nil {
		return exceptions.Internal("failed to delete flight", result.Error)
	}
	return nil
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

func (f flightRepository) DeleteIntermediateStops(flightID uint) error {
	result := f.db.Where("flight_id = ?", flightID).Delete(&models.IntermediateStop{})
	if result.Error != nil {
		return exceptions.Internal("failed to delete intermediate stops", result.Error)
	}
	return nil
}

func (f flightRepository) GetDB() *gorm.DB {
	return f.db
}

func (r *flightRepository) GetFlightsByDateRange(startDate, endDate time.Time) ([]*models.Flight, error) {
	var flights []*models.Flight
	result := r.db.
		Preload("Plane").
		Preload("DepartureAirport").
		Preload("ArrivalAirport").
		Preload("IntermediateStops").
		Preload("IntermediateStops.Airport").
		Where("departure_date_time BETWEEN ? AND ?", startDate, endDate).
		Find(&flights)
	if result.Error != nil {
		return nil, exceptions.Internal("failed to get flights by date range", result.Error)
	}
	return flights, nil
}
