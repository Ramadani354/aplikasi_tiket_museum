package models

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	ID      uint `json:"id"`
	IDAdmin uint `gorm:"column:id_admin" json:"-"`
	Harga   uint `json:"harga"`
	Kuota   uint `json:"kuota"`
}
