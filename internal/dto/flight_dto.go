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

type SeatClassInfo struct {
	ClassName   string `json:"class_name"`
	TotalSeats  int    `json:"total_seats"`
	BookedSeats int    `json:"booked_seats"`
	EmptySeats  int    `json:"empty_seats"`
}

type SeatInfo struct {
	SeatNumber string  `json:"seat_number"`
	ClassName  string  `json:"class_name"`
	IsBooked   bool    `json:"is_booked"`
	BookedBy   string  `json:"booked_by,omitempty"`
	Price      float64 `json:"price"`
}

type FlightResponse struct {
	FlightCode        string                `json:"flight_code"`
	DepartureAirport  string                `json:"departure_airport"`
	ArrivalAirport    string                `json:"arrival_airport"`
	Duration          int                   `json:"duration"`
	BasePrice         float64               `json:"base_price"`
	DepartureDateTime string                `json:"departure_date_time"`
	PlaneCode         string                `json:"plane_code"`
	IntermediateStop  []IntermediateStopDTO `json:"intermediate_stop"`
	EmptySeats        int                   `json:"empty_seats"`
	BookedSeats       int                   `json:"booked_seats"`
	TotalSeats        int                   `json:"total_seats"`
}

type FlightResponseDetailed struct {
	FlightCode        string                `json:"flight_code"`
	DepartureAirport  string                `json:"departure_airport"`
	ArrivalAirport    string                `json:"arrival_airport"`
	Duration          int                   `json:"duration"`
	BasePrice         float64               `json:"base_price"`
	DepartureDateTime string                `json:"departure_date_time"`
	PlaneCode         string                `json:"plane_code"`
	IntermediateStop  []IntermediateStopDTO `json:"intermediate_stop"`
	EmptySeats        int                   `json:"empty_seats"`
	BookedSeats       int                   `json:"booked_seats"`
	TotalSeats        int                   `json:"total_seats"`
	SeatClassInfo     []SeatClassInfo       `json:"seat_class_info"`
	Seats             []SeatInfo            `json:"seats"`
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

type FlightRevenueReport struct {
	FlightCode string  `json:"flightCode"`
	Tickets    int     `json:"tickets"`
	Revenue    float64 `json:"revenue"`
	Ratio      float64 `json:"ratio"` // Ratio of actual revenue to potential revenue
}

type MonthlyRevenueReport struct {
	Month        string                `json:"month"` // Format: "YYYY-MM"
	Flights      []FlightRevenueReport `json:"flights"`
	TotalRevenue float64               `json:"totalRevenue"`
	TotalTickets int                   `json:"totalTickets"`
	AverageRatio float64               `json:"averageRatio"`
}

type MonthlyRevenueSummary struct {
	Month       string  `json:"month"` // Format: "YYYY-MM"
	FlightCount int     `json:"flightCount"`
	Revenue     float64 `json:"revenue"`
	Ratio       float64 `json:"ratio"` // Average ratio across all flights
}

type YearlyRevenueReport struct {
	Year         string                  `json:"year"` // Format: "YYYY"
	Months       []MonthlyRevenueSummary `json:"months"`
	TotalRevenue float64                 `json:"totalRevenue"`
	TotalFlights int                     `json:"totalFlights"`
	AverageRatio float64                 `json:"averageRatio"`
}
