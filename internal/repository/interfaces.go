package repository

import (
	"errors"
	"fmt"
	"github.com/aprilboiz/flight-management/internal/models"
)

var (
	ErrNotFound        = errors.New("entity not found")
	ErrAlreadyExists   = errors.New("entity already exists")
	ErrFailedOperation = errors.New("repository operation failed")
	ErrInvalidArgument = errors.New("invalid repository argument")
)

type FlightRepository interface {
	GetAll() ([]*models.Flight, error)
	GetByID(id int) (*models.Flight, error)
	GetByCode(code string) (*models.Flight, error)
	Create(flight *models.Flight) (*models.Flight, error)
	Update(flight *models.Flight) (*models.Flight, error)
	Delete(flight *models.Flight) error
	CreateIntermediateStops(stops []*models.IntermediateStop) ([]*models.IntermediateStop, error)
}

type AirportRepository interface {
	GetAll() ([]*models.Airport, error)
	GetByCode(code string) (*models.Airport, error)
	GetByCodes(codes []string) (map[string]*models.Airport, error)
}

type PlaneRepository interface {
	GetByID(id uint) (*models.Plane, error)
	GetByCode(code string) (*models.Plane, error)
}

type TicketClassRepository interface {
	GetByName(name string) (*models.TicketClass, error)
	GetByNames(names []string) (map[string]*models.TicketClass, error)
}

type FlightCodeGenerator interface {
	Generate() (string, error)
}

func WrapError(baseErr error, msg string) error {
	return fmt.Errorf("%s: %w", msg, baseErr)
}
