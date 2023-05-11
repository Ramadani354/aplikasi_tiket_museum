package repositories

import (
	"github.com/Ramadani354/tiket_museum/app/models"
	"gorm.io/gorm"
)

type AdminRepository interface {
	GetAdminByEmail(email string) (*models.Admin, error)
	CreateAdmin(admin *models.Admin) error
	GetAdminByID(adminID uint) (*models.Admin, error)
	GetTicketQuota(adminID uint) (uint, error)
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

func (r *adminRepository) GetAdminByID(adminID uint) (*models.Admin, error) {
	admin := new(models.Admin)
	if err := r.db.First(admin, adminID).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *adminRepository) CreateAdmin(admin *models.Admin) error {
	return r.db.Create(admin).Error
}

func (r *adminRepository) GetTicketQuota(adminID uint) (uint, error) {
	var quota uint
	result := r.db.Model(&models.Ticket{}).Where("id_admin = ?", adminID).Select("SUM(kuota)").Scan(&quota)
	if result.Error != nil {
		return 0, result.Error
	}
	return quota, nil
}
