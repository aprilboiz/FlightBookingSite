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
	StopOrder    int    `json:"stop_order,omitempty"`
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
	EmptySeats        int                   `json:"empty_seats"`
	BookedSeats       int                   `json:"booked_seats"`
	TotalSeats        int                   `json:"total_seats"`
	IntermediateStop  []IntermediateStopDTO `json:"intermediate_stops"`
}

// FlightListResponse represents a flight in the list view
type FlightListResponse struct {
	FlightCode        string  `json:"flight_code"`
	PlaneCode         string  `json:"plane_code"`
	PlaneName         string  `json:"plane_name"`
	DepartureAirport  string  `json:"departure_airport"`
	DepartureCity     string  `json:"departure_city"`
	DepartureCountry  string  `json:"departure_country"`
	ArrivalAirport    string  `json:"arrival_airport"`
	ArrivalCity       string  `json:"arrival_city"`
	ArrivalCountry    string  `json:"arrival_country"`
	DepartureDateTime string  `json:"departure_date_time"`
	Duration          int     `json:"duration"`
	BasePrice         float64 `json:"base_price"`
	EmptySeats        int     `json:"empty_seats"`
	BookedSeats       int     `json:"booked_seats"`
	TotalSeats        int     `json:"total_seats"`
	HasStops          bool    `json:"has_stops"`
	StopCount         int     `json:"stop_count"`
}
