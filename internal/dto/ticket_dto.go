package dto

import "github.com/aprilboiz/flight-management/internal/models"

type TicketRequest struct {
	FlightCode  string             `json:"flight_code" binding:"required"`
	SeatNumber  string             `json:"seat_number" binding:"required"`
	FullName    string             `json:"full_name" binding:"required"`
	IDCard      string             `json:"id_card" binding:"required"`
	PhoneNumber string             `json:"phone_number" binding:"required"`
	Email       string             `json:"email" binding:"required,email"`
	BookingType models.BookingType `json:"booking_type" binding:"required,oneof=TICKET PLACE_ORDER"`
}

type TicketStatusUpdateRequest struct {
	Status models.TicketStatus `json:"status" binding:"required,oneof=ACTIVE CANCELLED EXPIRED"`
}

type TicketResponse struct {
	ID           uint                `json:"id"`
	FlightCode   string              `json:"flight_code"`
	SeatNumber   string              `json:"seat_number"`
	Price        float64             `json:"price"`
	FullName     string              `json:"full_name"`
	IDCard       string              `json:"id_card"`
	PhoneNumber  string              `json:"phone_number"`
	Email        string              `json:"email"`
	TicketStatus models.TicketStatus `json:"ticket_status"`
	BookingType  models.BookingType  `json:"booking_type"`
}

type TicketStatusesResponse struct {
	Statuses []models.TicketStatus `json:"statuses"`
}

type BookingTypesResponse struct {
	Types []models.BookingType `json:"types"`
}
