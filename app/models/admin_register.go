package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"-"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
