package services

import (
	"github.com/Ramadani354/tiket_museum/app/models"
	"github.com/Ramadani354/tiket_museum/app/repositories"
	"gorm.io/gorm"
)

type PrintService struct {
	printRepo repositories.PrintRepository
}

func NewPrintService(db *gorm.DB) *PrintService {
	return &PrintService{
		printRepo: repositories.NewPrintRepository(db),
	}
}

func (s *PrintService) CreatePrint(ticketID uint, tanggalCetak string) (*models.Print, error) {
	print := &models.Print{
		TicketID:     ticketID,
		TanggalCetak: tanggalCetak,
	}

	err := s.printRepo.CreatePrint(print)
	if err != nil {
		return nil, err
	}

	return print, nil
}
