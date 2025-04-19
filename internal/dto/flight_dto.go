package dto

type FlightRequest struct {
	DepartureAirport  string                `json:"departure_airport"`
	ArrivalAirport    string                `json:"arrival_airport"`
	Duration          int                   `json:"duration"`
	BasePrice         float64               `json:"base_price"`
	DepartureDateTime string                `json:"departure_date"`
	PlaneCode         string                `json:"plane_code"`
	IntermediateStop  []IntermediateStopDTO `json:"intermediate_stops"`
}

type IntermediateStopDTO struct {
	StopAirport  string `json:"stop_airport"`
	StopDuration int    `json:"stop_duration"`
	StopOrder    int    `json:"stop_order"`
	Note         string `json:"note"`
}

type FlightResponse struct {
	FlightCode        string                `json:"flight_code"`
	DepartureAirport  string                `json:"departure_airport"`
	ArrivalAirport    string                `json:"arrival_airport"`
	Duration          int                   `json:"duration"`
	BasePrice         float64               `json:"base_price"`
	DepartureDateTime string                `json:"departure_date"`
	PlaneCode         string                `json:"plane_code"`
	IntermediateStop  []IntermediateStopDTO `json:"intermediate_stop"`
}

type FlightResponseInList struct {
	FlightCode        string `json:"flight_code"`
	PlaneCode         string `json:"plane_code"`
	DepartureAirport  string `json:"departure_airport"`
	ArrivalAirport    string `json:"arrival_airport"`
	DepartureDateTime string `json:"departure_date"`
	EmptySeats        int    `json:"empty_seats"`
	BookedSeats       int    `json:"booked_seats"`
}
