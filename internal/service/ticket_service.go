package service

import (
	"fmt"
	"time"

	"github.com/aprilboiz/flight-management/internal/dto"
	"github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/internal/models"
	"github.com/aprilboiz/flight-management/internal/repository"
)

type ticketService struct {
	ticketRepo repository.TicketRepository
	flightRepo repository.FlightRepository
	planeRepo  repository.PlaneRepository
	paramRepo  repository.ParameterRepository
}

func (t *ticketService) GetAllTickets() ([]*dto.TicketResponse, error) {
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
			Price:        ticket.Price,
			FullName:     ticket.FullName,
			IDCard:       ticket.IDCard,
			PhoneNumber:  ticket.PhoneNumber,
			Email:        ticket.Email,
			TicketStatus: ticket.TicketStatus,
			BookingType:  ticket.BookingType,
		})
	}
	return tickets, nil
}

func (t *ticketService) GetTicketByID(id uint) (*dto.TicketResponse, error) {
	ticket, err := t.ticketRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.TicketResponse{
		ID:           ticket.ID,
		FlightCode:   ticket.Flight.FlightCode,
		SeatNumber:   ticket.Seat.SeatNumber,
		Price:        ticket.Price,
		FullName:     ticket.FullName,
		IDCard:       ticket.IDCard,
		PhoneNumber:  ticket.PhoneNumber,
		Email:        ticket.Email,
		TicketStatus: ticket.TicketStatus,
		BookingType:  ticket.BookingType,
	}, nil
}

func (t *ticketService) Create(ticket *dto.TicketRequest) (*dto.TicketResponse, error) {
	// 1. Validate flight exists and is not in the past
	flight, err := t.flightRepo.GetByCode(ticket.FlightCode)
	if err != nil {
		return nil, err
	}

	// Check if a flight is in the past
	if flight.DepartureDateTime.Before(time.Now()) {
		return nil, exceptions.BadRequestError("cannot book ticket for a past flight", nil)
	}

	// 2. Validate seat exists and is available
	seat, err := t.planeRepo.GetSeatByNumberAndPlaneCode(ticket.SeatNumber, flight.Plane.PlaneCode)
	if err != nil {
		return nil, err
	}

	// Check if a seat is already booked for this flight
	var existingTicket models.Ticket
	result := t.ticketRepo.GetDB().Where("flight_id = ? AND seat_id = ?", flight.ID, seat.ID).First(&existingTicket)
	if result.Error == nil {
		return nil, exceptions.BadRequestError("seat is already booked for this flight", nil)
	}

	// 3. Calculate ticket price based on seat class
	basePrice := flight.BasePrice
	ticketPrice := basePrice * seat.TicketClass.PricePercentage

	// 4. Validate booking type and timing
	bookingType := models.BookingTypeTicket
	if ticket.BookingType == models.BookingTypePlaceOrder {
		bookingType = models.BookingTypePlaceOrder

		// Get parameters for place order timing
		params, err := t.paramRepo.GetAllParams()
		if err != nil {
			return nil, err
		}

		// Check if place order is within the allowed time window
		daysBefore := time.Duration(params.LatestTicketPurchaseTime) * 24 * time.Hour
		deadline := flight.DepartureDateTime.Add(-daysBefore)
		if time.Now().After(deadline) {
			return nil, exceptions.BadRequestError(fmt.Sprintf("place orders must be made at least %d days before departure", params.LatestTicketPurchaseTime), nil)
		}
	}

	// 5. Create the ticket
	newTicket := &models.Ticket{
		FlightID:     flight.ID,
		SeatID:       seat.ID,
		Price:        ticketPrice,
		FullName:     ticket.FullName,
		IDCard:       ticket.IDCard,
		PhoneNumber:  ticket.PhoneNumber,
		Email:        ticket.Email,
		TicketStatus: models.TicketStatusActive,
		BookingType:  bookingType,
	}

	createdTicket, err := t.ticketRepo.Create(newTicket)
	if err != nil {
		return nil, err
	}

	// 6. Return the response
	return &dto.TicketResponse{
		ID:           createdTicket.ID,
		FlightCode:   flight.FlightCode,
		SeatNumber:   seat.SeatNumber,
		Price:        createdTicket.Price,
		FullName:     createdTicket.FullName,
		IDCard:       createdTicket.IDCard,
		PhoneNumber:  createdTicket.PhoneNumber,
		Email:        createdTicket.Email,
		TicketStatus: createdTicket.TicketStatus,
		BookingType:  createdTicket.BookingType,
	}, nil
}

func (t *ticketService) ConvertPlaceOrderToTicket(placeOrderID uint) (*dto.TicketResponse, error) {
	// 1. Get the place order
	placeOrder, err := t.ticketRepo.GetByID(placeOrderID)
	if err != nil {
		return nil, err
	}

	// 2. Validate it a place order
	if placeOrder.BookingType != models.BookingTypePlaceOrder {
		return nil, exceptions.BadRequestError("this is not a place order", nil)
	}

	// 3. Get the flight to check timing
	flight, err := t.flightRepo.GetByID(placeOrder.FlightID)
	if err != nil {
		return nil, err
	}

	// 4. Get parameters for timing validation
	params, err := t.paramRepo.GetAllParams()
	if err != nil {
		return nil, err
	}

	// 5. Check if conversion is within an allowed time window
	daysBefore := time.Duration(params.LatestTicketPurchaseTime) * 24 * time.Hour
	deadline := flight.DepartureDateTime.Add(-daysBefore)
	if time.Now().After(deadline) {
		return nil, exceptions.BadRequestError("cannot convert place order to ticket after the deadline", nil)
	}

	// 6. Update the booking type and status
	placeOrder.BookingType = models.BookingTypeTicket
	placeOrder.TicketStatus = models.TicketStatusActive

	updatedTicket, err := t.ticketRepo.Update(placeOrder)
	if err != nil {
		return nil, err
	}

	// 7. Return the response
	return &dto.TicketResponse{
		ID:           updatedTicket.ID,
		FlightCode:   flight.FlightCode,
		SeatNumber:   updatedTicket.Seat.SeatNumber,
		Price:        updatedTicket.Price,
		FullName:     updatedTicket.FullName,
		IDCard:       updatedTicket.IDCard,
		PhoneNumber:  updatedTicket.PhoneNumber,
		Email:        updatedTicket.Email,
		TicketStatus: updatedTicket.TicketStatus,
		BookingType:  updatedTicket.BookingType,
	}, nil
}

func (t *ticketService) CancelPlaceOrders(flightCode string) error {
	// Get the flight
	flight, err := t.flightRepo.GetByCode(flightCode)
	if err != nil {
		return err
	}

	// Get all place orders for this flight
	tickets, err := t.ticketRepo.GetByFlightID(flight.ID)
	if err != nil {
		return err
	}

	// Update each place to expire status
	for _, ticket := range tickets {
		if ticket.BookingType == models.BookingTypePlaceOrder {
			if err := t.ticketRepo.UpdateTicketStatus(ticket.ID, models.TicketStatusExpired); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *ticketService) UpdateTicketStatus(ticketID uint, newStatus models.TicketStatus) (*dto.TicketResponse, error) {
	// Get the ticket
	ticket, err := t.ticketRepo.GetByID(ticketID)
	if err != nil {
		return nil, err
	}

	// If trying to cancel, check if it's within the cancellation window
	if newStatus == models.TicketStatusCancelled {
		// Get the flight to check timing
		flight, err := t.flightRepo.GetByID(ticket.FlightID)
		if err != nil {
			return nil, err
		}

		// Get parameters for cancellation time
		params, err := t.paramRepo.GetAllParams()
		if err != nil {
			return nil, err
		}

		// Calculate cancellation deadline
		daysBefore := time.Duration(params.TicketCancellationTime) * 24 * time.Hour
		deadline := flight.DepartureDateTime.Add(-daysBefore)

		// Check if it's past the cancellation deadline
		if time.Now().After(deadline) {
			return nil, exceptions.BadRequestError(fmt.Sprintf("cannot cancel ticket after %d days before departure", params.TicketCancellationTime), nil)
		}

		// Check if the ticket is already cancelled
		if ticket.TicketStatus == models.TicketStatusCancelled {
			return nil, exceptions.BadRequestError("ticket is already cancelled", nil)
		}

		// Check if the ticket is already used
		if ticket.TicketStatus == models.TicketStatusUsed {
			return nil, exceptions.BadRequestError("cannot cancel a used ticket", nil)
		}
	}

	// Update ticket status
	if err := t.ticketRepo.UpdateTicketStatus(ticketID, newStatus); err != nil {
		return nil, err
	}

	// Get an updated ticket
	updatedTicket, err := t.ticketRepo.GetByID(ticketID)
	if err != nil {
		return nil, err
	}

	// Return response
	return &dto.TicketResponse{
		ID:           updatedTicket.ID,
		FlightCode:   updatedTicket.Flight.FlightCode,
		SeatNumber:   updatedTicket.Seat.SeatNumber,
		Price:        updatedTicket.Price,
		FullName:     updatedTicket.FullName,
		IDCard:       updatedTicket.IDCard,
		PhoneNumber:  updatedTicket.PhoneNumber,
		Email:        updatedTicket.Email,
		TicketStatus: updatedTicket.TicketStatus,
		BookingType:  updatedTicket.BookingType,
	}, nil
}

func (t *ticketService) DeleteTicket(id uint) error {
	// Validate ticket exists
	_, err := t.ticketRepo.GetByID(id)
	if err != nil {
		return err
	}

	// Delete the ticket
	return t.ticketRepo.Delete(id)
}

// GetTicketStatuses returns all available ticket statuses
func (t *ticketService) GetTicketStatuses() []models.TicketStatus {
	return []models.TicketStatus{
		models.TicketStatusActive,
		models.TicketStatusCancelled,
		models.TicketStatusUsed,
		models.TicketStatusExpired,
		models.TicketStatusRefunded,
	}
}

// GetBookingTypes returns all available booking types
func (t *ticketService) GetBookingTypes() []models.BookingType {
	return []models.BookingType{
		models.BookingTypeTicket,
		models.BookingTypePlaceOrder,
	}
}

func NewTicketService(ticketRepo repository.TicketRepository, flightRepo repository.FlightRepository, planeRepo repository.PlaneRepository, paramRepo repository.ParameterRepository) TicketService {
	return &ticketService{
		ticketRepo: ticketRepo,
		flightRepo: flightRepo,
		planeRepo:  planeRepo,
		paramRepo:  paramRepo,
	}
}
