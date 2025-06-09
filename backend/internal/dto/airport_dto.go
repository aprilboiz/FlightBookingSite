package dto

type AirportResponse struct {
	AirportCode string `json:"airport_code"`
	AirportName string `json:"airport_name"`
	CityName    string `json:"city_name"`
	CountryName string `json:"country_name"`
}
