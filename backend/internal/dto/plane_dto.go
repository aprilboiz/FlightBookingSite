package dto

type PlaneResponse struct {
	PlaneCode string `json:"plane_code"`
	PlaneName string `json:"plane_name"`
}

type PlaneResponseDetails struct {
	PlaneCode string         `json:"plane_code"`
	PlaneName string         `json:"plane_name"`
	Seats     []SeatResponse `json:"seats"`
}

type SeatResponse struct {
	SeatNumber  string `json:"seat_number"`
	TicketClass string `json:"ticket_class"`
}
