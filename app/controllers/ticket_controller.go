package controllers

import (
	"net/http"
	"strconv"

	"github.com/Ramadani354/tiket_museum/app/models"
	"github.com/Ramadani354/tiket_museum/app/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TicketController struct {
	ticketService *services.TicketService
}

func NewTicketController(db *gorm.DB) *TicketController {
	return &TicketController{
		ticketService: services.NewTicketService(db),
	}
}

func (c *TicketController) GetAllTickets(ctx echo.Context) error {
	tickets, err := c.ticketService.GetAllTickets()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to fetch tickets")
	}

	return ctx.JSON(http.StatusOK, tickets)
}

func (c *TicketController) CreateTicket(ctx echo.Context) error {
	payload := new(models.Ticket)
	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	ticket, err := c.ticketService.CreateTicket(payload)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to create ticket")
	}

	return ctx.JSON(http.StatusCreated, ticket)
}

func (c *TicketController) GetTicket(ctx echo.Context) error {
	ticketID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid ticket ID")
	}

	ticket, err := c.ticketService.GetTicket(uint(ticketID))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Ticket not found")
	}

	return ctx.JSON(http.StatusOK, ticket)
}

func (c *TicketController) UpdateTicket(ctx echo.Context) error {
	ticketID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid ticket ID")
	}

	payload := new(models.Ticket)
	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	ticket, err := c.ticketService.UpdateTicket(uint(ticketID), payload)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to update ticket")
	}

	return ctx.JSON(http.StatusOK, ticket)
}

func (c *TicketController) DeleteTicket(ctx echo.Context) error {
	ticketID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid ticket ID")
	}

	if err := c.ticketService.DeleteTicket(uint(ticketID)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to delete ticket")
	}

	return ctx.JSON(http.StatusOK, "Ticket deleted successfully")
}

// Tambahkan method GetTicketQuota di controllers/ticket_controller.go
func (c *TicketController) GetTicketQuota(ctx echo.Context) error {
	// Panggil repository yang sesuai untuk mengambil data kuota tiket dari database
	quota, err := c.ticketService.GetTicketQuota()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to get ticket quota")
	}

	return ctx.JSON(http.StatusOK, quota)
}
