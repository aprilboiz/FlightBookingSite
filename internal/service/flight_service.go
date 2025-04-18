package service

import (
	"github.com/aprilboiz/flight-management/internal/dto"
	"github.com/aprilboiz/flight-management/internal/models"
	"github.com/aprilboiz/flight-management/internal/repository"
	"time"
)

type flightService struct {
	flightRepo  repository.FlightRepository
	airportRepo repository.AirportRepository
	planeRepo   repository.PlaneRepository
}

func NewFlightService(flightRepo repository.FlightRepository) FlightService {
	return &flightService{flightRepo: flightRepo}
}

func (f flightService) GetAllFlights() ([]*dto.FlightResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (f flightService) GetFlightByID(flightID string) (*dto.FlightResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (f flightService) GetFlightByCode(flightCode string) (*dto.FlightResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (f flightService) Create(flightRequest *dto.FlightRequest) (*dto.FlightResponse, error) {
	flightCode := "Hello"
	plane, err := f.planeRepo.GetByCode(flightRequest.PlaneCode)
	if err != nil {
		return nil, err
	}
	departureAirport, err := f.airportRepo.GetByCode(flightRequest.DepartureAirport)
	if err != nil {
		return nil, err
	}
	arrivalAirport, err := f.airportRepo.GetByCode(flightRequest.ArrivalAirport)
	if err != nil {
		return nil, err
	}
	departureDateTime, err := time.Parse(time.DateTime, flightRequest.DepartureDateTime)
	if err != nil {
		return nil, err
	}

	newFlight := &models.Flight{
		FlightCode:         flightCode,
		PlaneID:            plane.ID,
		DepartureAirportID: departureAirport.ID,
		ArrivalAirportID:   arrivalAirport.ID,
		DepartureDateTime:  departureDateTime,
		FlightDuration:     flightRequest.Duration,
		BasePrice:          flightRequest.BasePrice,
	}
	createdFlight, err := f.flightRepo.Create(newFlight)
	if err != nil {
		return nil, err
	}

	intermediateStops := make([]*models.IntermediateStop, len(flightRequest.IntermediateStop))
	for i, stop := range flightRequest.IntermediateStop {
		airport, err := f.airportRepo.GetByCode(stop.StopAirport)
		if err != nil {
			return nil, err
		}
		intermediateStops[i] = &models.IntermediateStop{
			FlightID:     createdFlight.ID,
			AirportID:    airport.ID,
			StopDuration: stop.StopDuration,
			Note:         stop.Note,
		}
	}
	createdIntermediateStops, err := f.flightRepo.CreateIntermediateStops(intermediateStops)
	if err != nil {
		return nil, err
	}

	// Create the flight response DTO
	// Map the created intermediate stops to DTOs
	intermediateStopDTOs := make([]dto.IntermediateStopDTO, len(createdIntermediateStops))
	for i, stop := range createdIntermediateStops {
		intermediateStopDTOs[i] = dto.IntermediateStopDTO{
			StopAirport:  stop.Airport.AirportCode,
			StopDuration: stop.StopDuration,
			Note:         stop.Note,
		}
	}
	flightResponse := &dto.FlightResponse{
		FlightCode:        createdFlight.FlightCode,
		DepartureAirport:  createdFlight.DepartureAirport.AirportCode,
		ArrivalAirport:    createdFlight.ArrivalAirport.AirportCode,
		Duration:          createdFlight.FlightDuration,
		BasePrice:         createdFlight.BasePrice,
		DepartureDateTime: createdFlight.DepartureDateTime.Format(time.RFC3339),
		PlaneCode:         createdFlight.Plane.PlaneCode,
		IntermediateStop:  intermediateStopDTOs,
	}
	return flightResponse, nil
}

func (f flightService) Update(flightCode string, flightRequest *dto.FlightRequest) (*dto.FlightResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (f flightService) DeleteByCode(code string) error {
	//TODO implement me
	panic("implement me")
}
