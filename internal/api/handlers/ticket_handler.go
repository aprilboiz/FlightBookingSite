package handlers

import (
	"github.com/aprilboiz/flight-management/internal/dto"
	e "github.com/aprilboiz/flight-management/internal/exceptions"
	"github.com/aprilboiz/flight-management/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ticketHandler struct {
	ticketService service.TicketService
}

func (t ticketHandler) GetAllTickets(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t ticketHandler) GetTicketByID(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t ticketHandler) CreateTicket(c *gin.Context) {
	validatedModel, exists := c.Get("validatedModel")
	if !exists {
		_ = c.Error(e.NewAppError(e.INTERNAL_ERROR, "Cannot find validated model in context", nil))
		return
	}
	ticketRequest, ok := validatedModel.(*dto.TicketRequest)
	if !ok {
		_ = c.Error(e.NewAppError(e.INTERNAL_ERROR, "Cannot cast validated model to TicketRequest", nil))
		return
	}

	ticket, err := t.ticketService.Create(ticketRequest)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, ticket)
}

func (t ticketHandler) UpdateTicketStatus(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t ticketHandler) DeleteTicket(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewTicketHandler(ticketService service.TicketService) TicketHandler {
	return &ticketHandler{ticketService: ticketService}
}
