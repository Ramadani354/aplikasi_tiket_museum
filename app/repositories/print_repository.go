package repositories

import (
	"github.com/Ramadani354/tiket_museum/app/models"
	"gorm.io/gorm"
)

type PrintRepository interface {
	CreatePrint(print *models.Print) error
}

type printRepository struct {
	db *gorm.DB
}

func NewPrintRepository(db *gorm.DB) PrintRepository {
	return &printRepository{
		db: db,
	}
}

func (r *printRepository) CreatePrint(print *models.Print) error {
	return r.db.Create(print).Error
}
