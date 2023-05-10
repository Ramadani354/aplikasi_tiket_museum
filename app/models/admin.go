package models

type Admin struct {
	ID       uint   `json:"id"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
