package repository

import (
	"errors"
	"strconv"

	"github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/internal/models"
	"gorm.io/gorm"
)

type ticketRepository struct {
	db *gorm.DB
}

func (t ticketRepository) GetAll() ([]*models.Ticket, error) {
	var tickets []*models.Ticket
	result := t.db.
		Preload("Flight").
		Preload("Seat").
		Find(&tickets)
	if result.Error != nil {
		return nil, exceptions.Internal("failed to get all tickets", result.Error)
	}
	return tickets, nil
}

func (t ticketRepository) GetByID(id uint) (*models.Ticket, error) {
	var ticket models.Ticket
	result := t.db.
		Preload("Flight").
		Preload("Seat").
		Where("id = ?", id).
		First(&ticket)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, exceptions.NotFound("flight", strconv.Itoa(int(id)))
		}
		return nil, exceptions.Internal("failed to get flight by id", result.Error)
	}
	return &ticket, nil
}

func (t ticketRepository) Create(ticket *models.Ticket) (*models.Ticket, error) {
	result := t.db.Create(ticket)
	if result.Error != nil {
		return nil, exceptions.Internal("failed to create ticket", result.Error)
	}
	return ticket, nil
}

func (t ticketRepository) GetByFlightID(flightID uint) ([]*models.Ticket, error) {
	var tickets []*models.Ticket
	result := t.db.
		Preload("Flight").
		Preload("Seat").
		Where("flight_id = ?", flightID).
		Find(&tickets)
	if result.Error != nil {
		return nil, exceptions.Internal("failed to get tickets by flight ID", result.Error)
	}
	return tickets, nil
}

func (t ticketRepository) GetActiveTicketsByFlightID(flightID uint) ([]*models.Ticket, error) {
	var tickets []*models.Ticket
	result := t.db.
		Preload("Flight").
		Preload("Seat").
		Where("flight_id = ? AND ticket_status = ?", flightID, models.TicketStatusActive).
		Find(&tickets)
	if result.Error != nil {
		return nil, exceptions.Internal("failed to get active tickets by flight ID", result.Error)
	}
	return tickets, nil
}

func (t ticketRepository) UpdateTicketStatus(ticketID uint, status models.TicketStatus) error {
	result := t.db.Model(&models.Ticket{}).Where("id = ?", ticketID).Update("ticket_status", status)
	if result.Error != nil {
		return exceptions.Internal("failed to update ticket status", result.Error)
	}
	return nil
}

func (t ticketRepository) Update(ticket *models.Ticket) (*models.Ticket, error) {
	result := t.db.Save(ticket)
	if result.Error != nil {
		return nil, exceptions.Internal("failed to update ticket", result.Error)
	}
	return ticket, nil
}

func (t ticketRepository) GetDB() *gorm.DB {
	return t.db
}

func (t ticketRepository) Delete(id uint) error {
	result := t.db.Delete(&models.Ticket{}, id)
	if result.Error != nil {
		return exceptions.Internal("failed to delete ticket", result.Error)
	}
	return nil
}

func (r *ticketRepository) GetTicketsByFlightID(flightID uint) ([]*models.Ticket, error) {
	var tickets []*models.Ticket
	result := r.db.
		Where("flight_id = ?", flightID).
		Find(&tickets)
	if result.Error != nil {
		return nil, exceptions.Internal("failed to get tickets by flight ID", result.Error)
	}
	return tickets, nil
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db: db}
}
