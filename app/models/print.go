package models

import "gorm.io/gorm"

type Print struct {
	gorm.Model
	ID           uint   `gorm:"primary_key" json:"id"`
	IDTicket     uint   `gorm:"foreignKey:ID" json:"id_ticket"`
	TanggalCetak string `gorm:"column:tanggal_cetak" json:"tanggal_cetak"`
}
