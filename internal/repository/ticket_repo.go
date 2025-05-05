package repository

import (
	"errors"
	"github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/internal/models"
	"gorm.io/gorm"
	"strconv"
)

type ticketRepository struct {
	db *gorm.DB
}

func (t ticketRepository) GetAll() ([]*models.Ticket, error) {
	//TODO implement me
	panic("implement me")
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

func (t ticketRepository) UpdateTicketStatus(ticketId uint, newStatus string) (*models.Ticket, error) {
	//TODO implement me
	panic("implement me")
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db: db}
}
