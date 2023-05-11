package services

import (
	"errors"

	"github.com/Ramadani354/tiket_museum/app/models"
	"github.com/Ramadani354/tiket_museum/app/repositories"
	"github.com/Ramadani354/tiket_museum/utils"
	"gorm.io/gorm"
)

type AdminService struct {
	adminRepo repositories.AdminRepository
}

func NewAdminService(db *gorm.DB) *AdminService {
	return &AdminService{
		adminRepo: repositories.NewAdminRepository(db),
	}
}

func (s *AdminService) RegisterAdmin(admin *models.Admin) error {
	hashedPassword, err := utils.GenerateHashedPassword(admin.Password)
	if err != nil {
		return err
	}

	admin.Password = hashedPassword

	if err := s.adminRepo.CreateAdmin(admin); err != nil {
		return err
	}

	return nil
}

func (s *AdminService) LoginAdmin(email, password string) (*models.Admin, error) {
	admin, err := s.adminRepo.GetAdminByEmail(email)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(password, admin.Password) {
		return nil, errors.New("incorrect password")
	}

	return admin, nil
}

func (s *AdminService) GetTicketQuota(adminID uint) (uint, error) {
	quota, err := s.adminRepo.GetTicketQuota(adminID)
	if err != nil {
		return 0, err
	}

	return quota, nil
}
