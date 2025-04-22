package repository

import (
	"errors"
	"github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/internal/models"
	"gorm.io/gorm"
	"strconv"
)

type planeRepository struct {
	db *gorm.DB
}

func NewPlaneRepository(db *gorm.DB) PlaneRepository {
	return &planeRepository{db: db}
}

func (p planeRepository) GetByCode(code string) (*models.Plane, error) {
	var plane models.Plane
	result := p.db.Where("plane_code = ?", code).First(&plane)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, exceptions.NotFound("plane", code)
		}
		return nil, exceptions.Internal("failed to get plane by code", result.Error)
	}
	return &plane, nil
}

func (p planeRepository) GetByID(id uint) (*models.Plane, error) {
	var plane models.Plane
	result := p.db.Where("id = ?", id).First(&plane)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, exceptions.NotFound("plane", strconv.Itoa(int(id)))
		}
		return nil, exceptions.Internal("failed to get plane by id", result.Error)
	}
	return &plane, nil
}
