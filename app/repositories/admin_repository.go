package repositories

import (
	"github.com/Ramadani354/tiket_museum/app/models"
	"gorm.io/gorm"
)

type AdminRepository interface {
	GetAdminByEmail(email string) (*models.Admin, error)
	CreateAdmin(admin *models.Admin) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{
		db: db,
	}
}

func (r *adminRepository) GetAdminByEmail(email string) (*models.Admin, error) {
	admin := new(models.Admin)
	err := r.db.Where("email = ?", email).First(admin).Error
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *adminRepository) CreateAdmin(admin *models.Admin) error {
	return r.db.Create(admin).Error
}
