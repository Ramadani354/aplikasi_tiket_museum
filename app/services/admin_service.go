package services

import (
	"errors"

	"github.com/Ramadani354/tiket_museum/app/models"
	"github.com/Ramadani354/tiket_museum/app/repositories"
	"github.com/Ramadani354/tiket_museum/utils"
)

type AdminService struct {
	adminRepo repositories.AdminRepository
}

func NewAdminService() *AdminService {
	return &AdminService{
		adminRepo: repositories.NewAdminRepository(config.DB),
	}
}

func (s *AdminService) RegisterAdmin(admin *models.Admin) error {
	admin.Password = utils.HashPassword(admin.Password)
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
