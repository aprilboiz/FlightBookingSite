package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/aprilboiz/flight-management/internal/dto"
	"github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/internal/models"
	"github.com/aprilboiz/flight-management/internal/repository"
	"github.com/aprilboiz/flight-management/pkg/config"
	"github.com/aprilboiz/flight-management/pkg/database"
	"gorm.io/gorm"
)

type flightService struct {
	paramRepo   repository.ParameterRepository
	flightRepo  repository.FlightRepository
	airportRepo repository.AirportRepository
	planeRepo   repository.PlaneRepository
	ticketRepo  repository.TicketRepository
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

func NewFlightService(flightRepo repository.FlightRepository, airportRepo repository.AirportRepository, planeRepo repository.PlaneRepository, paramRepo repository.ParameterRepository, ticketRepo repository.TicketRepository) FlightService {

	if flightRepo == nil || airportRepo == nil || planeRepo == nil || paramRepo == nil || ticketRepo == nil {
		panic("Missing required repositories for flight service")
	}
	return &flightService{
		paramRepo:   paramRepo,
		flightRepo:  flightRepo,
		airportRepo: airportRepo,
		planeRepo:   planeRepo,
		ticketRepo:  ticketRepo,
	}
}

// Helper structs for seat counting
type PlaneSeatCount struct {
	PlaneID    uint  `gorm:"column:plane_id"`
	TotalSeats int64 `gorm:"column:total_seats"`
}

type FlightBookedCount struct {
	FlightID    uint  `gorm:"column:flight_id"`
	BookedSeats int64 `gorm:"column:booked_seats"`
}

// Helper function to get total seats for a plane
func (f flightService) getTotalSeatsForPlane(planeID uint) (int64, error) {
	var planeSeatCounts []PlaneSeatCount
	result := f.planeRepo.GetDB().Model(&models.Seat{}).
		Select("plane_id, COUNT(*) as total_seats").
		Where("plane_id = ?", planeID).
		Group("plane_id").
		Find(&planeSeatCounts)
	if result.Error != nil {
		return 0, exceptions.Internal("failed to count total seats", result.Error)
	}
	if len(planeSeatCounts) == 0 {
		return 0, nil
	}
	return planeSeatCounts[0].TotalSeats, nil
}

// Helper function to get booked seats for a flight
func (f flightService) getBookedSeatsForFlight(flightID uint) (int64, error) {
	var flightBookedCounts []FlightBookedCount
	result := f.ticketRepo.GetDB().Model(&models.Ticket{}).
		Select("flight_id, COUNT(*) as booked_seats").
		Where("flight_id = ? AND ticket_status = ?", flightID, models.TicketStatusActive).
		Group("flight_id").
		Find(&flightBookedCounts)
	if result.Error != nil {
		return 0, exceptions.Internal("failed to count booked seats", result.Error)
	}
	if len(flightBookedCounts) == 0 {
		return 0, nil
	}
	return flightBookedCounts[0].BookedSeats, nil
}

// Helper function to get seat counts for multiple flights
func (f flightService) getSeatCountsForFlights(flights []*models.Flight) (map[uint]int64, map[uint]int64, error) {
	// Get all flight IDs and plane IDs
	flightIDs := make([]uint, len(flights))
	planeIDs := make([]uint, len(flights))
	for i, flight := range flights {
		flightIDs[i] = flight.ID
		planeIDs[i] = flight.PlaneID
	}

	// Get total seats for each plane
	var planeSeatCounts []PlaneSeatCount
	result := f.planeRepo.GetDB().Model(&models.Seat{}).
		Select("plane_id, COUNT(*) as total_seats").
		Where("plane_id IN ?", planeIDs).
		Group("plane_id").
		Find(&planeSeatCounts)
	if result.Error != nil {
		return nil, nil, exceptions.Internal("failed to count total seats", result.Error)
	}

	// Convert to map for easier lookup
	planeSeats := make(map[uint]int64)
	for _, count := range planeSeatCounts {
		planeSeats[count.PlaneID] = count.TotalSeats
	}

	// Get booked seats for each flight
	var flightBookedCounts []FlightBookedCount
	result = f.ticketRepo.GetDB().Model(&models.Ticket{}).
		Select("flight_id, COUNT(*) as booked_seats").
		Where("flight_id IN ? AND ticket_status = ?", flightIDs, models.TicketStatusActive).
		Group("flight_id").
		Find(&flightBookedCounts)
	if result.Error != nil {
		return nil, nil, exceptions.Internal("failed to count booked seats", result.Error)
	}

	// Convert to map for easier lookup
	flightBookedSeats := make(map[uint]int64)
	for _, count := range flightBookedCounts {
		flightBookedSeats[count.FlightID] = count.BookedSeats
	}

	return planeSeats, flightBookedSeats, nil
}

func (f flightService) GetAllFlights() ([]*dto.FlightResponse, error) {
	// Get all flights with their relations
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

	// If there are no flights, return an empty list
	if len(flights) == 0 {
		return []*dto.FlightResponse{}, nil
	}

	// Get seat counts for all flights
	planeSeats, flightBookedSeats, err := f.getSeatCountsForFlights(flights)
	if err != nil {
		return nil, err
	}

	// Map flights to response
	flightResponses := make([]*dto.FlightResponse, len(flights))
	for i, flight := range flights {
		totalSeats := planeSeats[flight.PlaneID]
		bookedSeats := flightBookedSeats[flight.ID]
		emptySeats := totalSeats - bookedSeats

		intermediateStopDTOs := make([]dto.IntermediateStopDTO, len(flight.IntermediateStops))
		for j, stop := range flight.IntermediateStops {
			intermediateStopDTOs[j] = dto.IntermediateStopDTO{
				StopAirport:  stop.Airport.AirportCode,
				StopDuration: stop.StopDuration,
				StopOrder:    stop.StopOrder,
				Note:         stop.Note,
			}
		}

		flightResponses[i] = &dto.FlightResponse{
			FlightCode:        flight.FlightCode,
			DepartureAirport:  flight.DepartureAirport.AirportCode,
			ArrivalAirport:    flight.ArrivalAirport.AirportCode,
			Duration:          flight.FlightDuration,
			BasePrice:         flight.BasePrice,
			DepartureDateTime: flight.DepartureDateTime.Format(time.RFC3339),
			PlaneCode:         flight.Plane.PlaneCode,
			IntermediateStop:  intermediateStopDTOs,
			EmptySeats:        int(emptySeats),
			BookedSeats:       int(bookedSeats),
			TotalSeats:        int(totalSeats),
		}
	}

	return flightResponses, nil
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

	// Get total seats for the plane
	type PlaneSeatCount struct {
		PlaneID    uint  `gorm:"column:plane_id"`
		TotalSeats int64 `gorm:"column:total_seats"`
	}
	var planeSeatCounts []PlaneSeatCount
	result := f.planeRepo.GetDB().Model(&models.Seat{}).
		Select("plane_id, COUNT(*) as total_seats").
		Where("plane_id = ?", flight.PlaneID).
		Group("plane_id").
		Find(&planeSeatCounts)
	if result.Error != nil {
		return nil, exceptions.Internal("failed to count total seats", result.Error)
	}

	// Get booked seats for the flight
	type FlightBookedCount struct {
		FlightID    uint  `gorm:"column:flight_id"`
		BookedSeats int64 `gorm:"column:booked_seats"`
	}
	var flightBookedCounts []FlightBookedCount
	result = f.ticketRepo.GetDB().Model(&models.Ticket{}).
		Select("flight_id, COUNT(*) as booked_seats").
		Where("flight_id = ? AND ticket_status = ?", flight.ID, models.TicketStatusActive).
		Group("flight_id").
		Find(&flightBookedCounts)
	if result.Error != nil {
		return nil, exceptions.Internal("failed to count booked seats", result.Error)
	}

	// Calculate seat counts
	var totalSeats, bookedSeats int64
	if len(planeSeatCounts) > 0 {
		totalSeats = planeSeatCounts[0].TotalSeats
	}
	if len(flightBookedCounts) > 0 {
		bookedSeats = flightBookedCounts[0].BookedSeats
	}
	emptySeats := totalSeats - bookedSeats

	// Map intermediate stops
	intermediateStopDTOs := make([]dto.IntermediateStopDTO, len(flight.IntermediateStops))
	for i, stop := range flight.IntermediateStops {
		intermediateStopDTOs[i] = dto.IntermediateStopDTO{
			StopAirport:  stop.Airport.AirportCode,
			StopDuration: stop.StopDuration,
			StopOrder:    stop.StopOrder,
			Note:         stop.Note,
		}
	}

	return &dto.FlightResponse{
		FlightCode:        flight.FlightCode,
		DepartureAirport:  flight.DepartureAirport.AirportCode,
		ArrivalAirport:    flight.ArrivalAirport.AirportCode,
		Duration:          flight.FlightDuration,
		BasePrice:         flight.BasePrice,
		DepartureDateTime: flight.DepartureDateTime.Format(time.RFC3339),
		PlaneCode:         flight.Plane.PlaneCode,
		IntermediateStop:  intermediateStopDTOs,
		EmptySeats:        int(emptySeats),
		BookedSeats:       int(bookedSeats),
		TotalSeats:        int(totalSeats),
	}, nil
}

func (f flightService) Create(flightRequest *dto.FlightRequest) (*dto.FlightResponse, error) {
	// 1. Get parameters for validation
	params, err := f.paramRepo.GetAllParams()
	if err != nil {
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			return nil, appErr
		}
		return nil, exceptions.Internal("failed to get all params", err)
	}

	// 2. Validate basic flight requirements
	if flightRequest.DepartureAirport == flightRequest.ArrivalAirport {
		return nil, exceptions.BadRequest("departure and arrival airports cannot be the same", nil)
	}

	// 3. Validate flight duration
	if flightRequest.Duration < params.MinFlightDuration {
		return nil, exceptions.BadRequest(fmt.Sprintf("flight duration must be at least %d minutes", params.MinFlightDuration), nil)
	}

	// 4. Validate intermediate stops
	if len(flightRequest.IntermediateStop) > 0 {
		// Check maximum number of stops
		if len(flightRequest.IntermediateStop) > params.MaxIntermediateStops {
			return nil, exceptions.BadRequest(fmt.Sprintf("maximum number of intermediate stops is %d", params.MaxIntermediateStops), nil)
		}

		// Check for duplicate stops
		stopMap := make(map[string]bool)
		for _, stop := range flightRequest.IntermediateStop {
			// Check if stop is same as departure or arrival
			if stop.StopAirport == flightRequest.DepartureAirport {
				return nil, exceptions.BadRequest("intermediate stop cannot be the same as departure airport", nil)
			}
			if stop.StopAirport == flightRequest.ArrivalAirport {
				return nil, exceptions.BadRequest("intermediate stop cannot be the same as arrival airport", nil)
			}

			// Check for duplicate stops
			if stopMap[stop.StopAirport] {
				return nil, exceptions.BadRequest(fmt.Sprintf("duplicate intermediate stop airport: %s", stop.StopAirport), nil)
			}

			// Validate stop duration
			if stop.StopDuration < params.MinIntermediateStopDuration {
				return nil, exceptions.BadRequest(fmt.Sprintf("minimum intermediate stop duration is %d minutes", params.MinIntermediateStopDuration), nil)
			}
			if stop.StopDuration > params.MaxIntermediateStopDuration {
				return nil, exceptions.BadRequest(fmt.Sprintf("maximum intermediate stop duration is %d minutes", params.MaxIntermediateStopDuration), nil)
			}

			stopMap[stop.StopAirport] = true
		}

		// Validate stop order
		stopOrders := make(map[int]bool)
		for _, stop := range flightRequest.IntermediateStop {
			if stop.StopOrder < 1 {
				return nil, exceptions.BadRequest("stop order must be greater than 0", nil)
			}
			if stopOrders[stop.StopOrder] {
				return nil, exceptions.BadRequest(fmt.Sprintf("duplicate stop order: %d", stop.StopOrder), nil)
			}
			stopOrders[stop.StopOrder] = true
		}
	}

	// 5. Generate flight code
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

	// 6. Validate and get plane
	plane, err := f.planeRepo.GetByCode(flightRequest.PlaneCode)
	if err != nil {
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			return nil, appErr
		}
		return nil, exceptions.Internal("failed to get plane by code", err)
	}

	// 7. Validate and get airports
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

	// 8. Parse and validate departure date time
	loc, _ := time.LoadLocation(config.GetConfig().Database.Timezone)
	departureDateTime, err := time.ParseInLocation(time.DateTime, flightRequest.DepartureDateTime, loc)
	if err != nil {
		return nil, exceptions.BadRequest("invalid departure date time format", err)
	}

	// Validate departure is not in the past
	if departureDateTime.Before(time.Now()) {
		return nil, exceptions.BadRequest("departure date time cannot be in the past", nil)
	}

	// 9. Create the flight and intermediate stops in a transaction
	var createdFlight *models.Flight
	err = f.flightRepo.GetDB().Transaction(func(tx *gorm.DB) error {
		// Create the flight
		newFlight := &models.Flight{
			FlightCode:         flightCode,
			PlaneID:            plane.ID,
			DepartureAirportID: departureAirport.ID,
			ArrivalAirportID:   arrivalAirport.ID,
			DepartureDateTime:  departureDateTime,
			FlightDuration:     flightRequest.Duration,
			BasePrice:          flightRequest.BasePrice,
		}

		if err := tx.Create(newFlight).Error; err != nil {
			return exceptions.Internal("failed to create flight", err)
		}
		createdFlight = newFlight

		// Create intermediate stops if any
		if len(flightRequest.IntermediateStop) > 0 {
			intermediateStops := make([]*models.IntermediateStop, len(flightRequest.IntermediateStop))
			for i, stop := range flightRequest.IntermediateStop {
				airport, err := f.airportRepo.GetByCode(stop.StopAirport)
				if err != nil {
					var appErr *exceptions.AppError
					if errors.As(err, &appErr) {
						return appErr
					}
					return exceptions.Internal("failed to get intermediate airport by code", err)
				}
				intermediateStops[i] = &models.IntermediateStop{
					FlightID:     createdFlight.ID,
					AirportID:    airport.ID,
					StopDuration: stop.StopDuration,
					StopOrder:    stop.StopOrder,
					Note:         stop.Note,
				}
			}
			if err := tx.Create(&intermediateStops).Error; err != nil {
				return exceptions.Internal("failed to create intermediate stops", err)
			}
		}

		return nil
	})

	if err != nil {
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			return nil, appErr
		}
		return nil, err
	}

	// 10. Get the complete flight with all relations
	createdFlight, err = f.flightRepo.GetByCode(createdFlight.FlightCode)
	if err != nil {
		return nil, exceptions.Internal("failed to get created flight by code", err)
	}

	// Get seat counts
	totalSeats, err := f.getTotalSeatsForPlane(createdFlight.PlaneID)
	if err != nil {
		return nil, err
	}
	bookedSeats, err := f.getBookedSeatsForFlight(createdFlight.ID)
	if err != nil {
		return nil, err
	}
	emptySeats := totalSeats - bookedSeats

	// Map intermediate stops
	intermediateStopDTOs := make([]dto.IntermediateStopDTO, len(createdFlight.IntermediateStops))
	for i, stop := range createdFlight.IntermediateStops {
		intermediateStopDTOs[i] = dto.IntermediateStopDTO{
			StopAirport:  stop.Airport.AirportCode,
			StopDuration: stop.StopDuration,
			StopOrder:    stop.StopOrder,
			Note:         stop.Note,
		}
	}

	return &dto.FlightResponse{
		FlightCode:        createdFlight.FlightCode,
		DepartureAirport:  createdFlight.DepartureAirport.AirportCode,
		ArrivalAirport:    createdFlight.ArrivalAirport.AirportCode,
		Duration:          createdFlight.FlightDuration,
		BasePrice:         createdFlight.BasePrice,
		DepartureDateTime: createdFlight.DepartureDateTime.Format(time.RFC3339),
		PlaneCode:         createdFlight.Plane.PlaneCode,
		IntermediateStop:  intermediateStopDTOs,
		EmptySeats:        int(emptySeats),
		BookedSeats:       int(bookedSeats),
		TotalSeats:        int(totalSeats),
	}, nil
}

func (f flightService) Update(flightCode string, flightRequest *dto.FlightRequest) (*dto.FlightResponse, error) {
	// Get existing flight
	existingFlight, err := f.flightRepo.GetByCode(flightCode)
	if err != nil {
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			return nil, appErr
		}
		return nil, exceptions.Internal("failed to get flight by code", err)
	}

	// Get parameters for validation
	params, err := f.paramRepo.GetAllParams()
	if err != nil {
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			return nil, appErr
		}
		return nil, exceptions.Internal("failed to get all params", err)
	}

	// Validate basic flight requirements
	if flightRequest.DepartureAirport == flightRequest.ArrivalAirport {
		return nil, exceptions.BadRequest("departure and arrival airports cannot be the same", nil)
	}

	// Validate flight duration
	if flightRequest.Duration < params.MinFlightDuration {
		return nil, exceptions.BadRequest(fmt.Sprintf("flight duration must be at least %d minutes", params.MinFlightDuration), nil)
	}

	// Validate intermediate stops
	if len(flightRequest.IntermediateStop) > 0 {
		// Check maximum number of stops
		if len(flightRequest.IntermediateStop) > params.MaxIntermediateStops {
			return nil, exceptions.BadRequest(fmt.Sprintf("maximum number of intermediate stops is %d", params.MaxIntermediateStops), nil)
		}

		// Check for duplicate stops
		stopMap := make(map[string]bool)
		for _, stop := range flightRequest.IntermediateStop {
			// Check if stop is same as departure or arrival
			if stop.StopAirport == flightRequest.DepartureAirport {
				return nil, exceptions.BadRequest("intermediate stop cannot be the same as departure airport", nil)
			}
			if stop.StopAirport == flightRequest.ArrivalAirport {
				return nil, exceptions.BadRequest("intermediate stop cannot be the same as arrival airport", nil)
			}

			// Check for duplicate stops
			if stopMap[stop.StopAirport] {
				return nil, exceptions.BadRequest(fmt.Sprintf("duplicate intermediate stop airport: %s", stop.StopAirport), nil)
			}

			// Validate stop duration
			if stop.StopDuration < params.MinIntermediateStopDuration {
				return nil, exceptions.BadRequest(fmt.Sprintf("minimum intermediate stop duration is %d minutes", params.MinIntermediateStopDuration), nil)
			}
			if stop.StopDuration > params.MaxIntermediateStopDuration {
				return nil, exceptions.BadRequest(fmt.Sprintf("maximum intermediate stop duration is %d minutes", params.MaxIntermediateStopDuration), nil)
			}

			stopMap[stop.StopAirport] = true
		}
	}

	// Validate and get plane
	plane, err := f.planeRepo.GetByCode(flightRequest.PlaneCode)
	if err != nil {
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			return nil, appErr
		}
		return nil, exceptions.Internal("failed to get plane by code", err)
	}

	// Validate and get airports
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

	// Parse and validate departure date time
	loc, _ := time.LoadLocation(config.GetConfig().Database.Timezone)
	departureDateTime, err := time.ParseInLocation(time.DateTime, flightRequest.DepartureDateTime, loc)
	if err != nil {
		return nil, exceptions.BadRequest("invalid departure date time format", err)
	}

	// Update flight fields
	existingFlight.PlaneID = plane.ID
	existingFlight.DepartureAirportID = departureAirport.ID
	existingFlight.ArrivalAirportID = arrivalAirport.ID
	existingFlight.DepartureDateTime = departureDateTime
	existingFlight.FlightDuration = flightRequest.Duration
	existingFlight.BasePrice = flightRequest.BasePrice

	// Update the flight
	updatedFlight, err := f.flightRepo.Update(existingFlight)
	if err != nil {
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			return nil, appErr
		}
		return nil, exceptions.Internal("failed to update flight", err)
	}

	// Delete existing intermediate stops
	if err := f.flightRepo.DeleteIntermediateStops(updatedFlight.ID); err != nil {
		return nil, exceptions.Internal("failed to delete existing intermediate stops", err)
	}

	// Create new intermediate stops if any
	if len(flightRequest.IntermediateStop) > 0 {
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
				FlightID:     updatedFlight.ID,
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

	// Get the complete updated flight with all relations
	updatedFlight, err = f.flightRepo.GetByCode(updatedFlight.FlightCode)
	if err != nil {
		return nil, exceptions.Internal("failed to get updated flight by code", err)
	}

	// Get seat counts
	totalSeats, err := f.getTotalSeatsForPlane(updatedFlight.PlaneID)
	if err != nil {
		return nil, err
	}
	bookedSeats, err := f.getBookedSeatsForFlight(updatedFlight.ID)
	if err != nil {
		return nil, err
	}
	emptySeats := totalSeats - bookedSeats

	// Map intermediate stops
	intermediateStopDTOs := make([]dto.IntermediateStopDTO, len(updatedFlight.IntermediateStops))
	for i, stop := range updatedFlight.IntermediateStops {
		intermediateStopDTOs[i] = dto.IntermediateStopDTO{
			StopAirport:  stop.Airport.AirportCode,
			StopDuration: stop.StopDuration,
			StopOrder:    stop.StopOrder,
			Note:         stop.Note,
		}
	}

	return &dto.FlightResponse{
		FlightCode:        updatedFlight.FlightCode,
		DepartureAirport:  updatedFlight.DepartureAirport.AirportCode,
		ArrivalAirport:    updatedFlight.ArrivalAirport.AirportCode,
		Duration:          updatedFlight.FlightDuration,
		BasePrice:         updatedFlight.BasePrice,
		DepartureDateTime: updatedFlight.DepartureDateTime.Format(time.RFC3339),
		PlaneCode:         updatedFlight.Plane.PlaneCode,
		IntermediateStop:  intermediateStopDTOs,
		EmptySeats:        int(emptySeats),
		BookedSeats:       int(bookedSeats),
		TotalSeats:        int(totalSeats),
	}, nil
}

func (f flightService) Delete(code string) error {
	// Get the flight first to check if it exists
	flight, err := f.flightRepo.GetByCode(code)
	if err != nil {
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			return appErr
		}
		return exceptions.Internal("failed to get flight by code", err)
	}

	// Check if there are any active tickets for this flight
	activeTickets, err := f.ticketRepo.GetActiveTicketsByFlightID(flight.ID)
	if err != nil {
		return exceptions.Internal("failed to check active tickets", err)
	}
	if len(activeTickets) > 0 {
		return exceptions.BadRequest("cannot delete flight with active tickets", nil)
	}

	// Delete intermediate stops first
	if err := f.flightRepo.DeleteIntermediateStops(flight.ID); err != nil {
		return exceptions.Internal("failed to delete intermediate stops", err)
	}

	// Delete the flight
	if err := f.flightRepo.Delete(flight); err != nil {
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			return appErr
		}
		return exceptions.Internal("failed to delete flight", err)
	}

	return nil
}

func (f flightService) GetAllFlightsInList() ([]*dto.FlightListResponse, error) {
	// Get all flights with their relations
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

	// If there are no flights, return an empty list
	if len(flights) == 0 {
		return []*dto.FlightListResponse{}, nil
	}

	// Get seat counts for all flights
	planeSeats, flightBookedSeats, err := f.getSeatCountsForFlights(flights)
	if err != nil {
		return nil, err
	}

	// Map flights to response
	flightResponses := make([]*dto.FlightListResponse, len(flights))
	for i, flight := range flights {
		totalSeats := planeSeats[flight.PlaneID]
		bookedSeats := flightBookedSeats[flight.ID]
		emptySeats := totalSeats - bookedSeats

		flightResponses[i] = &dto.FlightListResponse{
			FlightCode:        flight.FlightCode,
			PlaneCode:         flight.Plane.PlaneCode,
			PlaneName:         flight.Plane.PlaneName,
			DepartureAirport:  flight.DepartureAirport.AirportCode,
			DepartureCity:     flight.DepartureAirport.CityName,
			DepartureCountry:  flight.DepartureAirport.CountryName,
			ArrivalAirport:    flight.ArrivalAirport.AirportCode,
			ArrivalCity:       flight.ArrivalAirport.CityName,
			ArrivalCountry:    flight.ArrivalAirport.CountryName,
			DepartureDateTime: flight.DepartureDateTime.Format(time.RFC3339),
			Duration:          flight.FlightDuration,
			BasePrice:         flight.BasePrice,
			EmptySeats:        int(emptySeats),
			BookedSeats:       int(bookedSeats),
			TotalSeats:        int(totalSeats),
			HasStops:          len(flight.IntermediateStops) > 0,
			StopCount:         len(flight.IntermediateStops),
		}
	}

	return flightResponses, nil
}
