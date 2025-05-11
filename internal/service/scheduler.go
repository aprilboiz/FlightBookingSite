package service

import (
	"time"

	"github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/internal/models"
	"github.com/aprilboiz/flight-management/internal/repository"
	"go.uber.org/zap"
)

type SchedulerService struct {
	ticketRepo repository.TicketRepository
	flightRepo repository.FlightRepository
	logger     *zap.Logger
}

func NewSchedulerService(ticketRepo repository.TicketRepository, flightRepo repository.FlightRepository) *SchedulerService {
	return &SchedulerService{
		ticketRepo: ticketRepo,
		flightRepo: flightRepo,
		logger:     zap.L(),
	}
}

func (s *SchedulerService) StartPlaceOrderCancellationJob() {
	// Run every hour
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for range ticker.C {
			if err := s.cancelExpiredPlaceOrders(); err != nil {
				s.logger.Error("Error cancelling expired place orders", zap.Error(err))
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
		return exceptions.Internal("failed to get place orders", result.Error)
	}

	// Cancel each place order
	for _, ticket := range tickets {
		ticket.TicketStatus = models.TicketStatusExpired
		if _, err := s.ticketRepo.Update(&ticket); err != nil {
			s.logger.Error("Error cancelling place order",
				zap.Uint("ticketID", ticket.ID),
				zap.Uint("flightID", ticket.FlightID),
				zap.Error(err))
			continue
		}
		s.logger.Debug("Expired place order",
			zap.Uint("ticketID", ticket.ID),
			zap.Uint("flightID", ticket.FlightID))
	}

	return nil
}
