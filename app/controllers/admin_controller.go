package controllers

import (
	"net/http"

	"github.com/Ramadani354/tiket_museum/app/models"
	"github.com/Ramadani354/tiket_museum/app/services"
	"github.com/Ramadani354/tiket_museum/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AdminController struct {
	adminService *services.AdminService
}

func NewAdminController(db *gorm.DB) *AdminController {
	return &AdminController{
		adminService: services.NewAdminService(db),
	}
}

func (c *AdminController) Register(ctx echo.Context) error {
	payload := new(models.Admin)
	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	if err := c.adminService.RegisterAdmin(payload); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to register admin")
	}

	return ctx.JSON(http.StatusOK, payload)
}

func (c *AdminController) Login(ctx echo.Context) error {
	payload := new(models.Login)
	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	admin, err := c.adminService.LoginAdmin(payload.Email, payload.Password)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, "Invalid email or password")
	}

	token, err := utils.GenerateJWTToken(admin.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to generate token")
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
