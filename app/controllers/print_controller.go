package controllers

import (
	"net/http"

	"github.com/Ramadani354/tiket_museum/app/services"
	"github.com/labstack/echo/v4"
)

type PrintController struct {
	printService services.PrintService
}

func NewPrintController(printService services.PrintService) *PrintController {
	return &PrintController{
		printService: printService,
	}
}

func (c *PrintController) CreatePrint(ctx echo.Context) error {
	// Buat struct untuk menampung data dari permintaan
	type request struct {
		TicketID     uint   `json:"ticket_id"`
		TanggalCetak string `json:"tanggal_cetak"`
	}

	// Bind data dari permintaan ke struct request
	req := new(request)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Permintaan tidak valid",
		})
	}

	// Panggil service untuk membuat cetakan tiket
	print, err := c.printService.CreatePrint(req.TicketID, req.TanggalCetak)
	if err != nil {
		// Tangani kesalahan jika terjadi kesalahan dalam membuat cetakan tiket
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Gagal membuat cetakan tiket",
		})
	}

	return ctx.JSON(http.StatusOK, print)
}
