package service

import (
	"github.com/aprilboiz/flight-management/internal/dto"
	"github.com/aprilboiz/flight-management/internal/models"
	"github.com/aprilboiz/flight-management/internal/repository"
)

type ticketService struct {
	ticketRepo repository.TicketRepository
	flightRepo repository.FlightRepository
	planeRepo  repository.PlaneRepository
}

func (t ticketService) GetAllTickets() ([]*dto.TicketResponse, error) {
	allTickets, err := t.ticketRepo.GetAll()
	tickets := make([]*dto.TicketResponse, 0)
	if err != nil {
		return nil, err
	}
	for _, ticket := range allTickets {
		tickets = append(tickets, &dto.TicketResponse{
			ID:           ticket.ID,
			FlightCode:   ticket.Flight.FlightCode,
			SeatNumber:   ticket.Seat.SeatNumber,
			TicketPrice:  ticket.Price,
			FullName:     ticket.FullName,
			IDCard:       ticket.IDCard,
			PhoneNumber:  ticket.PhoneNumber,
			Email:        ticket.Email,
			FlightStatus: ticket.FlightStatus,
		})
	}
	return tickets, nil
}

func (t ticketService) GetTicketByID(id uint) (*dto.TicketResponse, error) {
	ticket, err := t.ticketRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.TicketResponse{
		ID:           ticket.ID,
		FlightCode:   ticket.Flight.FlightCode,
		SeatNumber:   ticket.Seat.SeatNumber,
		TicketPrice:  ticket.Price,
		FullName:     ticket.FullName,
		IDCard:       ticket.IDCard,
		PhoneNumber:  ticket.PhoneNumber,
		Email:        ticket.Email,
		FlightStatus: ticket.FlightStatus,
	}, nil
}

func (t ticketService) Create(ticket *dto.TicketRequest) (*dto.TicketResponse, error) {
	flight, err := t.flightRepo.GetByCode(ticket.FlightCode)
	if err != nil {
		return nil, err
	}
	seat, err := t.planeRepo.GetSeatByNumberAndPlaneCode(ticket.SeatNumber, flight.Plane.PlaneCode)
	if err != nil {
		return nil, err
	}
	newTicket := &models.Ticket{
		FlightID:     flight.ID,
		SeatID:       seat.ID,
		Price:        ticket.TicketPrice,
		FullName:     ticket.FullName,
		IDCard:       ticket.IDCard,
		PhoneNumber:  ticket.PhoneNumber,
		Email:        ticket.Email,
		FlightStatus: ticket.FlightStatus,
	}

	createdTicket, err := t.ticketRepo.Create(newTicket)
	if err != nil {
		return nil, err
	}
	return &dto.TicketResponse{
		ID:           createdTicket.ID,
		FlightCode:   flight.FlightCode,
		SeatNumber:   seat.SeatNumber,
		TicketPrice:  createdTicket.Price,
		FullName:     createdTicket.FullName,
		IDCard:       createdTicket.IDCard,
		PhoneNumber:  createdTicket.PhoneNumber,
		Email:        createdTicket.Email,
		FlightStatus: createdTicket.FlightStatus,
	}, nil
}

func (t ticketService) UpdateTicketStatus(ticketId uint, newStatus string) (*dto.TicketResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewTicketService(ticketRepo repository.TicketRepository, flightRepo repository.FlightRepository, planeRepo repository.PlaneRepository) TicketService {
	return &ticketService{
		ticketRepo: ticketRepo,
		flightRepo: flightRepo,
		planeRepo:  planeRepo}
}
