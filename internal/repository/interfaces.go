package repository

import (
	"time"

	"github.com/aprilboiz/flight-management/internal/models"
	"gorm.io/gorm"
)

type FlightRepository interface {
	GetAll() ([]*models.Flight, error)
	GetByID(id uint) (*models.Flight, error)
	GetByCode(code string) (*models.Flight, error)
	Create(flight *models.Flight) (*models.Flight, error)
	Update(flight *models.Flight) (*models.Flight, error)
	Delete(flight *models.Flight) error
	CreateIntermediateStops(stops []*models.IntermediateStop) ([]*models.IntermediateStop, error)
	DeleteIntermediateStops(flightID uint) error
	GetDB() *gorm.DB
	GetFlightsByDateRange(startDate, endDate time.Time) ([]*models.Flight, error)
}

type AirportRepository interface {
	GetAll() ([]*models.Airport, error)
	GetByCode(code string) (*models.Airport, error)
	GetByCodes(codes []string) (map[string]*models.Airport, error)
	GetDB() *gorm.DB
}

type PlaneRepository interface {
	GetAll() ([]*models.Plane, error)
	GetByID(id uint) (*models.Plane, error)
	GetByCode(code string) (*models.Plane, error)
	GetSeatByNumberAndPlaneCode(seatNumber, planeCode string) (*models.Seat, error)
	GetDB() *gorm.DB
}

type TicketClassRepository interface {
	GetByName(name string) (*models.TicketClass, error)
	GetByNames(names []string) (map[string]*models.TicketClass, error)
	GetDB() *gorm.DB
}

type ParameterRepository interface {
	GetAllParams() (*models.Parameter, error)
	UpdateParams(params *models.Parameter) (*models.Parameter, error)
	GetDB() *gorm.DB
}

type FlightCodeGenerator interface {
	Generate() (string, error)
}

type TicketRepository interface {
	GetAll() ([]*models.Ticket, error)
	GetByID(id uint) (*models.Ticket, error)
	GetByFlightID(flightID uint) ([]*models.Ticket, error)
	GetActiveTicketsByFlightID(flightID uint) ([]*models.Ticket, error)
	Create(ticket *models.Ticket) (*models.Ticket, error)
	Update(ticket *models.Ticket) (*models.Ticket, error)
	UpdateTicketStatus(ticketID uint, status string) error
	Delete(id uint) error
	GetDB() *gorm.DB
	GetTicketsByFlightID(flightID uint) ([]*models.Ticket, error)
}
