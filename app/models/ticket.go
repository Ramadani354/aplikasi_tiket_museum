package models

type Ticket struct {
	ID      uint `json:"id"`
	IDAdmin uint `json:"id_admin"`
	Harga   uint `json:"harga"`
	Kuota   uint `json:"kuota"`
}
