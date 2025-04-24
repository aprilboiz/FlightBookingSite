package dto

type TicketRequest struct {
	FlightCode  string `json:"flight_code"`
	FullName    string `json:"full_name"`
	IDCard      string `json:"id_card"`
	PhoneNumber string `json:"phone_number"`
}
