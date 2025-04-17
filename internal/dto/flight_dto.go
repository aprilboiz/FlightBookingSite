package dto

type FlightRequest struct {
	BasePrice         float64             `json:"base_price"`
	DepartureAirport  string              `json:"departure"`
	ArrivalAirport    string              `json:"arrival"`
	DepartureDateTime string              `json:"departure_date"`
	Duration          int                 `json:"duration"`
	IntermediateStop  IntermediateStopDTO `json:"intermediate_stop"`
}

type IntermediateStopDTO struct {
	StopAirport  string `json:"stop_airport"`
	StopDuration int    `json:"stop_duration"`
	Note         string `json:"note"`
}
