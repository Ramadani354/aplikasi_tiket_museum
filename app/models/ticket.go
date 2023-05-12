package models

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	ID           uint      `gorm:"primaryKey" json:"id"`
	IDAdmin      uint      `gorm:"column:id_admin" json:"id_admin"`
	Harga        uint      `json:"harga"`
	Kuota        uint      `json:"kuota"`
	TanggalCetak time.Time `gorm:"column:tanggal_cetak; default:null" json:"tanggal_cetak"`
}
