package service

import (
	"github.com/aprilboiz/flight-management/internal/dto"
	"github.com/aprilboiz/flight-management/internal/models"
)

type FlightService interface {
	Create(flight *dto.FlightRequest) (*dto.FlightResponse, error)
	GetAllFlights() ([]*dto.FlightResponse, error)
	GetAllFlightsInList() ([]*dto.FlightListResponse, error)
	GetFlightByCode(flightCode string) (*dto.FlightResponseDetailed, error)
	Update(code string, flight *dto.FlightRequest) (*dto.FlightResponse, error)
	Delete(code string) error
	GetMonthlyRevenueReport(year int, month int) (*dto.MonthlyRevenueReport, error)
	GetYearlyRevenueReport(year int) (*dto.YearlyRevenueReport, error)
}

type AirportService interface {
	GetAllAirports() ([]*dto.AirportResponse, error)
	GetAirportByCode(code string) (*dto.AirportResponse, error)
	GetAirportsByCodes(codes []string) (map[string]*dto.AirportResponse, error)
}

type PlaneService interface {
	GetAllPlanes() ([]*dto.PlaneResponse, error)
	GetPlaneByCode(code string) (*dto.PlaneResponseDetails, error)
}

type ParameterService interface {
	GetAllParams() (*models.Parameter, error)
	UpdateParams(params *models.Parameter) (*models.Parameter, error)
}

type FlightCodeGenerator interface {
	Generate() (string, error)
}

type TicketService interface {
	Create(ticket *dto.TicketRequest) (*dto.TicketResponse, error)
	GetAllTickets() ([]*dto.TicketResponse, error)
	GetTicketByID(id uint) (*dto.TicketResponse, error)
	UpdateTicketStatus(ticketId uint, newStatus models.TicketStatus) (*dto.TicketResponse, error)
	DeleteTicket(id uint) error
	CancelPlaceOrders(flightCode string) error
	GetTicketStatuses() []models.TicketStatus
	GetBookingTypes() []models.BookingType
}

type UserService interface {
	Register(req dto.RegisterRequest) (*dto.AuthResponse, error)
	Login(req dto.LoginRequest) (*dto.AuthResponse, error)
}
