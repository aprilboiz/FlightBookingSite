package service

import (
	"github.com/aprilboiz/flight-management/internal/models"
	"github.com/aprilboiz/flight-management/internal/repository"
)

func NewParamService(paramRepo repository.ParameterRepository) ParameterService {
	return &paramService{
		paramRepo: paramRepo,
	}
}

type paramService struct {
	paramRepo repository.ParameterRepository
}

func (p paramService) GetAllParams() (*models.Parameter, error) {
	return p.paramRepo.GetAllParams()
}

func (p paramService) UpdateParams(params *models.Parameter) (*models.Parameter, error) {
	return p.paramRepo.UpdateParams(params)
}
