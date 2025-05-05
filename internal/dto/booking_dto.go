package dto

type TicketRequest struct {
	FlightCode   string  `json:"flight_code"`
	SeatNumber   string  `json:"seat_number"`
	FullName     string  `json:"full_name"`
	IDCard       string  `json:"id_card"`
	PhoneNumber  string  `json:"phone_number"`
	TicketPrice  float64 `json:"ticket_price"`
	Email        string  `json:"email"`
	FlightStatus string  `json:"flight_status"`
}

type TicketResponse struct {
	ID           uint    `json:"id"`
	FlightCode   string  `json:"flight_code"`
	SeatNumber   string  `json:"seat_number"`
	FullName     string  `json:"full_name"`
	IDCard       string  `json:"id_card"`
	PhoneNumber  string  `json:"phone_number"`
	TicketPrice  float64 `json:"ticket_price"`
	Email        string  `json:"email"`
	FlightStatus string  `json:"flight_status"`
}
