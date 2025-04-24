package service

import (
	"errors"
	"fmt"
	"github.com/aprilboiz/flight-management/internal/dto"
	"github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/internal/models"
	"github.com/aprilboiz/flight-management/internal/repository"
	"github.com/aprilboiz/flight-management/pkg/config"
	"github.com/aprilboiz/flight-management/pkg/database"
	"time"
)

type flightService struct {
	paramRepo   repository.ParameterRepository
	flightRepo  repository.FlightRepository
	airportRepo repository.AirportRepository
	planeRepo   repository.PlaneRepository
}

type flightCodeGenerator struct {
}

func NewFlightCodeGenerator() FlightCodeGenerator {
	return &flightCodeGenerator{}
}

func (g *flightCodeGenerator) Generate() (string, error) {
	nextID, err := database.PeekUpcomingFlightId()
	if err != nil {
		return "", exceptions.Internal("failed to get next ID", err)
	}

	// Format the ID to have leading zeros up to 4 digits
	formattedID := fmt.Sprintf("%04d", nextID)

	// Prefix with company name
	flightCode := "RuaAirline" + formattedID

	return flightCode, nil
}

func NewFlightService(flightRepo repository.FlightRepository, airportRepo repository.AirportRepository, planeRepo repository.PlaneRepository, paramRepo repository.ParameterRepository) FlightService {

	if flightRepo == nil || airportRepo == nil || planeRepo == nil || paramRepo == nil {
		panic("Missing required repositories for flight service")
	}
	return &flightService{
		paramRepo:   paramRepo,
		flightRepo:  flightRepo,
		airportRepo: airportRepo,
		planeRepo:   planeRepo,
	}
}

func (f flightService) GetAllFlights() ([]*dto.FlightResponse, error) {
	flights, err := f.flightRepo.GetAll()
	if err != nil {
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			return nil, &exceptions.AppError{
				Code:       appErr.Code,
				Message:    fmt.Sprintf("Error retrieving flights: %s", appErr.Message),
				StatusCode: appErr.StatusCode,
				Err:        appErr.Err,
			}
		}
		return nil, exceptions.Internal("Unexpected error retrieving flights", err)
	}
	flightResponses := make([]*dto.FlightResponse, len(flights))
	for i, flight := range flights {
		flightResponses[i] = &dto.FlightResponse{
			FlightCode:        flight.FlightCode,
			DepartureAirport:  flight.DepartureAirport.AirportCode,
			ArrivalAirport:    flight.ArrivalAirport.AirportCode,
			Duration:          flight.FlightDuration,
			BasePrice:         flight.BasePrice,
			DepartureDateTime: flight.DepartureDateTime.Format(time.RFC3339),
			PlaneCode:         flight.Plane.PlaneCode,
			IntermediateStop:  make([]dto.IntermediateStopDTO, len(flight.IntermediateStops)),
		}
		for j, stop := range flight.IntermediateStops {
			flightResponses[i].IntermediateStop[j] = dto.IntermediateStopDTO{
				StopAirport:  stop.Airport.AirportCode,
				StopDuration: stop.StopDuration,
				StopOrder:    stop.StopOrder,
				Note:         stop.Note,
			}
		}
	}
	return flightResponses, nil
}

func (f flightService) GetFlightByID(flightID int) (*dto.FlightResponse, error) {
	flight, err := f.flightRepo.GetByID(flightID)
	if err != nil {
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			return nil, &exceptions.AppError{
				Code:       appErr.Code,
				Message:    fmt.Sprintf("Error retrieving flight with ID '%d': %s", flightID, appErr.Message),
				StatusCode: appErr.StatusCode,
				Err:        appErr.Err,
			}
		}
		return nil, exceptions.Internal("Unexpected error retrieving flight", err)
	}
	flightResponse := &dto.FlightResponse{
		FlightCode:        flight.FlightCode,
		DepartureAirport:  flight.DepartureAirport.AirportCode,
		ArrivalAirport:    flight.ArrivalAirport.AirportCode,
		Duration:          flight.FlightDuration,
		BasePrice:         flight.BasePrice,
		DepartureDateTime: flight.DepartureDateTime.Format(time.RFC3339),
		PlaneCode:         flight.Plane.PlaneCode,
		IntermediateStop:  make([]dto.IntermediateStopDTO, len(flight.IntermediateStops)),
	}
	for i, stop := range flight.IntermediateStops {
		flightResponse.IntermediateStop[i] = dto.IntermediateStopDTO{
			StopAirport:  stop.Airport.AirportCode,
			StopDuration: stop.StopDuration,
			StopOrder:    stop.StopOrder,
			Note:         stop.Note,
		}
	}
	return flightResponse, nil
}

func (f flightService) GetFlightByCode(flightCode string) (*dto.FlightResponse, error) {
	flight, err := f.flightRepo.GetByCode(flightCode)
	if err != nil {
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			return nil, &exceptions.AppError{
				Code:       appErr.Code,
				Message:    fmt.Sprintf("Error retrieving flight '%s': %s", flightCode, appErr.Message),
				StatusCode: appErr.StatusCode,
				Err:        appErr.Err,
			}
		}
		return nil, exceptions.Internal("Unexpected error retrieving flight", err)
	}
	flightResponse := &dto.FlightResponse{
		FlightCode:        flight.FlightCode,
		DepartureAirport:  flight.DepartureAirport.AirportCode,
		ArrivalAirport:    flight.ArrivalAirport.AirportCode,
		Duration:          flight.FlightDuration,
		BasePrice:         flight.BasePrice,
		DepartureDateTime: flight.DepartureDateTime.Format(time.RFC3339),
		PlaneCode:         flight.Plane.PlaneCode,
		IntermediateStop:  make([]dto.IntermediateStopDTO, len(flight.IntermediateStops)),
	}
	for i, stop := range flight.IntermediateStops {
		flightResponse.IntermediateStop[i] = dto.IntermediateStopDTO{
			StopAirport:  stop.Airport.AirportCode,
			StopDuration: stop.StopDuration,
			StopOrder:    stop.StopOrder,
			Note:         stop.Note,
		}
	}
	return flightResponse, nil
}

func (f flightService) Create(flightRequest *dto.FlightRequest) (*dto.FlightResponse, error) {
	params, err := f.paramRepo.GetAllParams()
	if err != nil {
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			return nil, appErr
		}
		return nil, exceptions.Internal("failed to get all params", err)
	}

	if flightRequest.DepartureAirport == flightRequest.ArrivalAirport {
		return nil, exceptions.BadRequest("departure and arrival airports cannot be the same", nil)
	}

	if len(flightRequest.IntermediateStop) > 0 {
		stopMap := make(map[string]bool)
		if len(flightRequest.IntermediateStop) > params.MaxIntermediateStops {
			return nil, exceptions.BadRequest(fmt.Sprintf("the max intermediate stops is %v", params.MaxIntermediateStops), nil)
		}
		for _, stop := range flightRequest.IntermediateStop {
			if stopMap[stop.StopAirport] {
				return nil, exceptions.BadRequest(fmt.Sprintf("the stop airport '%s' is duplicated", stop.StopAirport), nil)
			}

			if stop.StopDuration < params.MinIntermediateStopDuration || stop.StopDuration > params.MaxIntermediateStopDuration {
				return nil, exceptions.BadRequest(fmt.Sprintf("the stop duration must be in range [%v, %v]", params.MinIntermediateStopDuration, params.MaxIntermediateStopDuration), nil)
			}
			stopMap[stop.StopAirport] = true
		}
	}

	if flightRequest.Duration < params.MinFlightDuration {
		return nil, exceptions.BadRequest(fmt.Sprintf("the min flight duration is %v", params.MinFlightDuration), nil)
	}

	flightCode, err := NewFlightCodeGenerator().Generate()
	if err != nil {
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			return nil, &exceptions.AppError{
				Code:       appErr.Code,
				Message:    fmt.Sprintf("Error generating flight code: %s", appErr.Message),
				StatusCode: appErr.StatusCode,
				Err:        appErr.Err,
			}
		}
		return nil, exceptions.Internal("failed to generate flight code", err)
	}
	plane, err := f.planeRepo.GetByCode(flightRequest.PlaneCode)
	if err != nil {
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			return nil, appErr
		}
		return nil, exceptions.Internal("failed to get plane by code", err)
	}
	departureAirport, err := f.airportRepo.GetByCode(flightRequest.DepartureAirport)
	if err != nil {
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			return nil, appErr
		}
		return nil, exceptions.Internal("failed to get departure airport by code", err)
	}
	arrivalAirport, err := f.airportRepo.GetByCode(flightRequest.ArrivalAirport)
	if err != nil {
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			return nil, appErr
		}
		return nil, exceptions.Internal("failed to get arrival airport by code", err)
	}
	loc, _ := time.LoadLocation(config.GetConfig().Database.Timezone)
	departureDateTime, err := time.ParseInLocation(time.DateTime, flightRequest.DepartureDateTime, loc)
	if err != nil {
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			return nil, appErr
		}
		return nil, exceptions.Internal("failed to parse departure date time", err)
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
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			return nil, appErr
		}
		return nil, exceptions.Internal("failed to create flight", err)
	}

	if len(flightRequest.IntermediateStop) != 0 {
		intermediateStops := make([]*models.IntermediateStop, len(flightRequest.IntermediateStop))
		for i, stop := range flightRequest.IntermediateStop {
			airport, err := f.airportRepo.GetByCode(stop.StopAirport)
			if err != nil {
				var appErr *exceptions.AppError
				if errors.As(err, &appErr) {
					return nil, appErr
				}
				return nil, exceptions.Internal("failed to get intermediate airport by code", err)
			}
			intermediateStops[i] = &models.IntermediateStop{
				FlightID:     createdFlight.ID,
				AirportID:    airport.ID,
				StopDuration: stop.StopDuration,
				StopOrder:    stop.StopOrder,
				Note:         stop.Note,
			}
		}
		_, err = f.flightRepo.CreateIntermediateStops(intermediateStops)
		if err != nil {
			var appErr *exceptions.AppError
			if errors.As(err, &appErr) {
				return nil, appErr
			}
			return nil, exceptions.Internal("failed to create intermediate stops", err)
		}
	}

	// CREATE THE FLIGHT RESPONSE DTO
	createdFlight, err = f.flightRepo.GetByCode(createdFlight.FlightCode)
	if err != nil {
		return nil, exceptions.Internal("failed to get created flight by code", err)
	}
	// Map the created intermediate stops to DTOs
	intermediateStopDTOs := make([]dto.IntermediateStopDTO, len(createdFlight.IntermediateStops))
	for i, stop := range createdFlight.IntermediateStops {
		intermediateStopDTOs[i] = dto.IntermediateStopDTO{
			StopAirport:  stop.Airport.AirportCode,
			StopDuration: stop.StopDuration,
			StopOrder:    stop.StopOrder,
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
