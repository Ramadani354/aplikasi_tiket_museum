package models

type Admin struct {
	ID       uint   `gorm:"primaryKey" json:"-"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
