package handlers

import (
	"net/http"
	"strconv"

	"github.com/aprilboiz/flight-management/internal/dto"
	e "github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/internal/service"
	"github.com/gin-gonic/gin"
)

type ticketHandler struct {
	ticketService service.TicketService
}

// GetAllTickets godoc
//	@Summary		Get all tickets
//	@Description	Retrieve a list of all tickets in the system
//	@Tags			tickets
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		dto.TicketResponse
//	@Failure		500	{object}	exceptions.AppError
//	@Router			/tickets [get]
func (t *ticketHandler) GetAllTickets(c *gin.Context) {
	tickets, err := t.ticketService.GetAllTickets()
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, tickets)
}

// GetTicketByID godoc
//	@Summary		Get ticket by ID
//	@Description	Retrieve a specific ticket by its ID
//	@Tags			tickets
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Ticket ID"
//	@Success		200	{object}	dto.TicketResponse
//	@Failure		404	{object}	exceptions.AppError
//	@Failure		500	{object}	exceptions.AppError
//	@Router			/tickets/{id} [get]
func (t *ticketHandler) GetTicketByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		_ = c.Error(e.NewAppError(e.BadRequest, "Invalid ticket ID format", err))
		return
	}

	ticket, err := t.ticketService.GetTicketByID(uint(id))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, ticket)
}

// CreateTicket godoc
//	@Summary		Create a new ticket
//	@Description	Create a new ticket with the provided information
//	@Tags			tickets
//	@Accept			json
//	@Produce		json
//	@Param			ticket	body		dto.TicketRequest	true	"Ticket information"
//	@Success		201		{object}	dto.TicketResponse
//	@Failure		400		{object}	exceptions.AppError
//	@Failure		500		{object}	exceptions.AppError
//	@Router			/tickets [post]
func (t *ticketHandler) CreateTicket(c *gin.Context) {
	validatedModel, exists := c.Get("validatedModel")
	if !exists {
		_ = c.Error(e.NewAppError(e.INTERNAL, "Cannot find validated model in context", nil))
		return
	}
	ticketRequest, ok := validatedModel.(*dto.TicketRequest)
	if !ok {
		_ = c.Error(e.NewAppError(e.INTERNAL, "Cannot cast validated model to TicketRequest", nil))
		return
	}

	ticket, err := t.ticketService.Create(ticketRequest)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, ticket)
}

// UpdateTicketStatus godoc
//
//	@Summary		Update ticket status
//	@Description	Update the status of a specific ticket
//	@Tags			tickets
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int								true	"Ticket ID"
//	@Param			status	body		dto.TicketStatusUpdateRequest	true	"New ticket status"
//	@Success		200		{object}	dto.TicketResponse
//	@Failure		400		{object}	exceptions.AppError
//	@Failure		404		{object}	exceptions.AppError
//	@Failure		500		{object}	exceptions.AppError
//	@Router			/tickets/{id}/status [put]
func (t *ticketHandler) UpdateTicketStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		_ = c.Error(e.NewAppError(e.BadRequest, "Invalid ticket ID format", err))
		return
	}

	validatedModel, exists := c.Get("validatedModel")
	if !exists {
		_ = c.Error(e.NewAppError(e.INTERNAL, "Cannot find validated model in context", nil))
		return
	}
	statusRequest, ok := validatedModel.(*dto.TicketStatusUpdateRequest)
	if !ok {
		_ = c.Error(e.NewAppError(e.INTERNAL, "Cannot cast validated model to TicketStatusUpdateRequest", nil))
		return
	}

	ticket, err := t.ticketService.UpdateTicketStatus(uint(id), statusRequest.Status)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, ticket)
}

// DeleteTicket godoc
//	@Summary		Delete a ticket
//	@Description	Delete a specific ticket by its ID
//	@Tags			tickets
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"Ticket ID"
//	@Success		204	"No Content"
//	@Failure		404	{object}	exceptions.AppError
//	@Failure		500	{object}	exceptions.AppError
//	@Router			/tickets/{id} [delete]
func (t *ticketHandler) DeleteTicket(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		_ = c.Error(e.NewAppError(e.BadRequest, "Invalid ticket ID format", err))
		return
	}

	if err := t.ticketService.DeleteTicket(uint(id)); err != nil {
		_ = c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}

// GetTicketStatuses godoc
//
//	@Summary		Get all available ticket statuses
//	@Description	Get a list of all possible ticket statuses
//	@Tags			tickets
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.TicketStatusesResponse
//	@Router			/tickets/statuses [get]
func (t *ticketHandler) GetTicketStatuses(c *gin.Context) {
	statuses := t.ticketService.GetTicketStatuses()
	c.JSON(http.StatusOK, dto.TicketStatusesResponse{
		Statuses: statuses,
	})
}

// GetBookingTypes godoc
//
//	@Summary		Get all available booking types
//	@Description	Get a list of all possible booking types
//	@Tags			tickets
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.BookingTypesResponse
//	@Router			/tickets/booking-types [get]
func (t *ticketHandler) GetBookingTypes(c *gin.Context) {
	types := t.ticketService.GetBookingTypes()
	c.JSON(http.StatusOK, dto.BookingTypesResponse{
		Types: types,
	})
}

func NewTicketHandler(ticketService service.TicketService) TicketHandler {
	return &ticketHandler{ticketService: ticketService}
}
