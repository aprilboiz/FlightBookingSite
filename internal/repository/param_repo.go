package repository

import (
	"github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/internal/models"
	"gorm.io/gorm"
)

func NewParameterRepository(db *gorm.DB) ParameterRepository {
	return &parameterRepository{db: db}
}

type parameterRepository struct {
	db *gorm.DB
}

func (p parameterRepository) GetAllParams() (*models.Parameter, error) {
	var params models.Parameter
	result := p.db.First(&params)
	if result.Error != nil {
		return nil, exceptions.Internal("failed to get all params", result.Error)
	}
	return &params, nil
}

func (p parameterRepository) UpdateParams(params *models.Parameter) (*models.Parameter, error) {
	var oldParams models.Parameter
	result := p.db.First(&oldParams)
	if result.Error != nil {
		return nil, exceptions.Internal("failed to get old params", result.Error)
	}
	result = p.db.Model(&oldParams).Updates(params)
	if result.Error != nil {
		return nil, exceptions.Internal("failed to update params", result.Error)
	}
	return &oldParams, nil
}

func (p parameterRepository) GetDB() *gorm.DB {
	return p.db
}
