package service

import (
	"github.com/aprilboiz/flight-management/internal/dto"
	"github.com/aprilboiz/flight-management/internal/repository"
)

func NewPlaneService(planeRepo repository.PlaneRepository) PlaneService {
	if planeRepo == nil {
		panic("Missing required repositories for plane service")
	}
	return &planeService{planeRepo: planeRepo}
}

type planeService struct {
	planeRepo repository.PlaneRepository
}

func (p planeService) GetAllPlanes() ([]*dto.PlaneResponse, error) {
	planes, err := p.planeRepo.GetAll()
	if err != nil {
		return nil, err
	}
	planeResponses := make([]*dto.PlaneResponse, len(planes))
	for i, plane := range planes {
		planeResponses[i] = &dto.PlaneResponse{
			PlaneCode: plane.PlaneCode,
			PlaneName: plane.PlaneName,
		}
	}
	return planeResponses, nil
}

func (p planeService) GetPlaneByCode(code string) (*dto.PlaneResponseDetails, error) {
	plane, err := p.planeRepo.GetByCode(code)
	if err != nil {
		return nil, err
	}

	planeResponse := &dto.PlaneResponseDetails{
		PlaneCode: plane.PlaneCode,
		PlaneName: plane.PlaneName,
		Seats:     make([]dto.SeatResponse, 0),
	}
	for _, seat := range plane.Seats {
		planeResponse.Seats = append(planeResponse.Seats, dto.SeatResponse{
			SeatNumber:  seat.SeatNumber,
			TicketClass: seat.TicketClass.TicketClassName,
		})
	}
	return planeResponse, nil
}
