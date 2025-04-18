package repository

import (
	"github.com/aprilboiz/flight-management/internal/models"
	"gorm.io/gorm"
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
		return nil, result.Error
	}
	return &plane, nil
}

func (p planeRepository) GetByID(id uint) (*models.Plane, error) {
	//TODO implement me
	panic("implement me")
}
