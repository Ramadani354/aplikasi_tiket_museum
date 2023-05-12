package services

import (
	"time"

	"github.com/Ramadani354/tiket_museum/app/models"
	"github.com/Ramadani354/tiket_museum/app/repositories"
	"gorm.io/gorm"
)

type PrintService struct {
	printRepo  repositories.PrintRepository
	ticketRepo repositories.TicketRepository
}

func NewPrintService(db *gorm.DB) *PrintService {
	return &PrintService{
		printRepo:  repositories.NewPrintRepository(db),
		ticketRepo: repositories.NewTicketRepository(db),
	}
}

func (s *PrintService) CreatePrint(ticketID uint) (*models.Print, error) {
	ticket, err := s.ticketRepo.GetTicketByID(ticketID)
	if err != nil {
		return nil, err
	}

	ticket.Kuota--

	err = s.ticketRepo.UpdateTicketKuota(ticketID, ticket.Kuota)
	if err != nil {
		return nil, err
	}

	print := &models.Print{
		TicketID:     ticket.ID,
		TanggalCetak: time.Now(),
	}

	err = s.printRepo.CreatePrint(print)
	if err != nil {
		return nil, err
	}

	return print, nil
}
