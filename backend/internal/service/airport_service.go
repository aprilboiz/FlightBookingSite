package service

import (
	"github.com/aprilboiz/flight-management/internal/dto"
	"github.com/aprilboiz/flight-management/internal/repository"
)

type airportService struct {
	airportRepo repository.AirportRepository
}

func NewAirportService(airportRepo repository.AirportRepository) AirportService {
	if airportRepo == nil {
		panic("Missing required repositories for airport service")
	}
	return &airportService{airportRepo: airportRepo}
}

func (a airportService) GetAllAirports() ([]*dto.AirportResponse, error) {
	airports, err := a.airportRepo.GetAll()
	if err != nil {
		return nil, err
	}

	airportResponses := make([]*dto.AirportResponse, len(airports))
	for i, airport := range airports {
		airportResponses[i] = &dto.AirportResponse{
			AirportCode: airport.AirportCode,
			AirportName: airport.AirportName,
			CityName:    airport.CityName,
			CountryName: airport.CountryName,
		}
	}
	return airportResponses, nil
}

func (a airportService) GetAirportByCode(code string) (*dto.AirportResponse, error) {
	airport, err := a.airportRepo.GetByCode(code)
	if err != nil {
		return nil, err
	}

	airportResponse := &dto.AirportResponse{
		AirportCode: airport.AirportCode,
		AirportName: airport.AirportName,
		CityName:    airport.CityName,
		CountryName: airport.CountryName,
	}
	return airportResponse, nil
}

func (a airportService) GetAirportsByCodes(codes []string) (map[string]*dto.AirportResponse, error) {
	airports, err := a.airportRepo.GetByCodes(codes)
	if err != nil {
		return nil, err
	}
	airportResponses := make(map[string]*dto.AirportResponse, len(airports))
	for _, airport := range airports {
		airportResponses[airport.AirportCode] = &dto.AirportResponse{
			AirportCode: airport.AirportCode,
			AirportName: airport.AirportName,
			CityName:    airport.CityName,
			CountryName: airport.CountryName,
		}
	}
	return airportResponses, nil
}
