package models

import (
	"time"

	"gorm.io/gorm"
)

type Print struct {
	gorm.Model
	ID           uint      `gorm:"primaryKey" json:"-"`
	TicketID     uint      `gorm:"column:id_ticket" json:"id_ticket"`
	TanggalCetak time.Time `gorm:"column:tanggal_cetak" json:"tanggal_cetak"`
}
