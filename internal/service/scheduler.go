package service

import (
	"log"
	"time"

	"github.com/aprilboiz/flight-management/internal/models"
	"github.com/aprilboiz/flight-management/internal/repository"
)

type SchedulerService struct {
	ticketRepo repository.TicketRepository
	flightRepo repository.FlightRepository
}

func NewSchedulerService(ticketRepo repository.TicketRepository, flightRepo repository.FlightRepository) *SchedulerService {
	return &SchedulerService{
		ticketRepo: ticketRepo,
		flightRepo: flightRepo,
	}
}

func (s *SchedulerService) StartPlaceOrderCancellationJob() {
	// Run every hour
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for range ticker.C {
			if err := s.cancelExpiredPlaceOrders(); err != nil {
				log.Printf("Error cancelling expired place orders: %v", err)
			}
		}
	}()
}

func (s *SchedulerService) cancelExpiredPlaceOrders() error {
	// Get all flights that are departing within the next 24 hours
	now := time.Now()
	next24Hours := now.Add(24 * time.Hour)

	// Find all place orders for flights departing within next 24 hours
	var tickets []models.Ticket
	result := s.ticketRepo.GetDB().
		Joins("JOIN flights ON flights.id = tickets.flight_id").
		Where("tickets.booking_type = ? AND flights.departure_date_time BETWEEN ? AND ?",
			models.BookingTypePlaceOrder, now, next24Hours).
		Find(&tickets)

	if result.Error != nil {
		return result.Error
	}

	// Cancel each place order
	for _, ticket := range tickets {
		ticket.TicketStatus = models.TicketStatusExpired
		if _, err := s.ticketRepo.Update(&ticket); err != nil {
			log.Printf("Error cancelling place order %d: %v", ticket.ID, err)
			continue
		}
		log.Printf("Expired place order %d for flight %d", ticket.ID, ticket.FlightID)
	}

	return nil
}
