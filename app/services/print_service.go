package services

import (
	"github.com/Ramadani354/tiket_museum/app/models"
	"github.com/Ramadani354/tiket_museum/app/repositories"
)

type PrintService struct {
	printRepo repositories.PrintRepository
}

func NewPrintService(repo repositories.PrintRepository) PrintService {
	return PrintService{
		printRepo: repo,
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
