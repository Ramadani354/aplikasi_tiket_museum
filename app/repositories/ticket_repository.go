package repositories

import (
	"github.com/Ramadani354/tiket_museum/app/models"
	"gorm.io/gorm"
)

type TicketRepository interface {
	GetAllTickets() ([]*models.Ticket, error)
	GetTicketByID(ticketID uint) (*models.Ticket, error)
	CreateTicket(ticket *models.Ticket) error
	UpdateTicket(ticket *models.Ticket) error
	DeleteTicket(ticket *models.Ticket) error
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{
		db: db,
	}
}

func (r *ticketRepository) GetAllTickets() ([]*models.Ticket, error) {
	var tickets []*models.Ticket
	err := r.db.Find(&tickets).Error
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (r *ticketRepository) GetTicketByID(ticketID uint) (*models.Ticket, error) {
	ticket := new(models.Ticket)
	err := r.db.First(ticket, ticketID).Error
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (r *ticketRepository) CreateTicket(ticket *models.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r *ticketRepository) UpdateTicket(ticket *models.Ticket) error {
	return r.db.Save(ticket).Error
}

func (r *ticketRepository) DeleteTicket(ticket *models.Ticket) error {
	return r.db.Delete(ticket).Error
}
