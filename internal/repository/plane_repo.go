package repository

import (
	"errors"
	"strconv"

	"github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/internal/models"
	"gorm.io/gorm"
)

func NewPlaneRepository(db *gorm.DB) PlaneRepository {
	return &planeRepository{db: db}
}

type planeRepository struct {
	db *gorm.DB
}

func (p planeRepository) GetSeatByNumberAndPlaneCode(seatNumber, planeCode string) (*models.Seat, error) {
	plane, err := p.GetByCode(planeCode)
	if err != nil {
		return nil, err
	}

	var seat models.Seat
	result := p.db.Where("seat_number = ? AND plane_id = ?", seatNumber, plane.ID).
		Preload("TicketClass").
		First(&seat)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, exceptions.NotFoundError("seat", seatNumber)
		}
		return nil, exceptions.InternalError("failed to get seat by number and plane code", result.Error)
	}
	return &seat, nil
}

func (p planeRepository) GetAll() ([]*models.Plane, error) {
	planes := make([]*models.Plane, 0)

	if err := p.db.Find(&planes).Error; err != nil {
		return nil, exceptions.InternalError("failed to get all planes", err)
	}

	return planes, nil
}

func (p planeRepository) GetByCode(code string) (*models.Plane, error) {
	var plane models.Plane
	result := p.db.Where("plane_code = ?", code).
		Preload("Seats.TicketClass").
		First(&plane)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, exceptions.NotFoundError("plane", code)
		}
		return nil, exceptions.InternalError("failed to get plane by code", result.Error)
	}
	return &plane, nil
}

func (p planeRepository) GetByID(id uint) (*models.Plane, error) {
	var plane models.Plane
	result := p.db.Where("id = ?", id).First(&plane)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, exceptions.NotFoundError("plane", strconv.Itoa(int(id)))
		}
		return nil, exceptions.InternalError("failed to get plane by id", result.Error)
	}
	return &plane, nil
}

func (p planeRepository) GetDB() *gorm.DB {
	return p.db
}
