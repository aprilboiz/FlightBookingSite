package handlers

import (
	"github.com/gin-gonic/gin"
)

type FlightHandler interface {
	GetAllFlightsInList(c *gin.Context)
	GetAllFlights(c *gin.Context)
	GetFlightByCode(c *gin.Context)
	CreateFlight(c *gin.Context)
	UpdateFlight(c *gin.Context)
	DeleteFlightByCode(c *gin.Context)
	GetMonthlyRevenueReport(c *gin.Context)
	GetYearlyRevenueReport(c *gin.Context)
	GetRevenueReport(c *gin.Context)
}

type PlaneHandler interface {
	GetAllPlanes(c *gin.Context)
	GetPlaneByCode(c *gin.Context)
}

type AirportHandler interface {
	GetAllAirports(c *gin.Context)
	GetAirportByCode(c *gin.Context)
}

type ParameterHandler interface {
	GetAllParameters(c *gin.Context)
	UpdateParameters(c *gin.Context)
}

type TicketHandler interface {
	GetAllTickets(c *gin.Context)
	GetTicketByID(c *gin.Context)
	CreateTicket(c *gin.Context)
	UpdateTicketStatus(c *gin.Context)
	DeleteTicket(c *gin.Context)
	GetTicketStatuses(c *gin.Context)
	GetBookingTypes(c *gin.Context)
}

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}
