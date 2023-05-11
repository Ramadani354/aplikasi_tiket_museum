package services

import (
	"github.com/Ramadani354/tiket_museum/app/models"
	"github.com/Ramadani354/tiket_museum/app/repositories"
	"gorm.io/gorm"
)

type TicketService struct {
	ticketRepo repositories.TicketRepository
}

func NewTicketService(db *gorm.DB) *TicketService {
	return &TicketService{
		ticketRepo: repositories.NewTicketRepository(db),
	}
}

func (s *TicketService) GetAllTickets() ([]*models.Ticket, error) {
	tickets, err := s.ticketRepo.GetAllTickets()
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (s *TicketService) CreateTicket(payload *models.Ticket) (*models.Ticket, error) {
	ticket := &models.Ticket{
		IDAdmin: payload.IDAdmin,
		Harga:   payload.Harga,
		Kuota:   payload.Kuota,
	}

	err := s.ticketRepo.CreateTicket(ticket)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (s *TicketService) GetTicket(ticketID uint) (*models.Ticket, error) {
	ticket, err := s.ticketRepo.GetTicketByID(ticketID)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (s *TicketService) UpdateTicket(ticketID uint, payload *models.Ticket) (*models.Ticket, error) {
	ticket, err := s.ticketRepo.GetTicketByID(ticketID)
	if err != nil {
		return nil, err
	}

	ticket.Harga = payload.Harga
	ticket.Kuota = payload.Kuota

	err = s.ticketRepo.UpdateTicket(ticket)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (s *TicketService) DeleteTicket(ticketID uint) error {
	ticket, err := s.ticketRepo.GetTicketByID(ticketID)
	if err != nil {
		return err
	}

	err = s.ticketRepo.DeleteTicket(ticket)
	if err != nil {
		return err
	}

	return nil
}

func (s *TicketService) GetTicketQuota() (uint, error) {
	return s.ticketRepo.GetTicketQuota()
}
