package repository

import (
	"github.com/aprilboiz/flight-management/internal/models"
)

type FlightRepository interface {
	GetAll() ([]*models.Flight, error)
	GetByID(id uint) (*models.Flight, error)
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
	GetAll() ([]*models.Plane, error)
	GetByID(id uint) (*models.Plane, error)
	GetByCode(code string) (*models.Plane, error)
	GetSeatByNumberAndPlaneCode(seatNumber, planeCode string) (*models.Seat, error)
}

type TicketClassRepository interface {
	GetByName(name string) (*models.TicketClass, error)
	GetByNames(names []string) (map[string]*models.TicketClass, error)
}

type ParameterRepository interface {
	GetAllParams() (*models.Parameter, error)
	UpdateParams(params *models.Parameter) (*models.Parameter, error)
}

type FlightCodeGenerator interface {
	Generate() (string, error)
}

type TicketRepository interface {
	GetAll() ([]*models.Ticket, error)
	GetByID(id uint) (*models.Ticket, error)
	Create(ticket *models.Ticket) (*models.Ticket, error)
	UpdateTicketStatus(ticketId uint, newStatus string) (*models.Ticket, error)
}
