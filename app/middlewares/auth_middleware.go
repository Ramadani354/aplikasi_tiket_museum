package middlewares

import (
	"net/http"
	"strings"

	"github.com/Ramadani354/tiket_museum/app/repositories"
	"github.com/Ramadani354/tiket_museum/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		claims, err := utils.VerifyJWTToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}

		userID := uint(claims["user_id"].(float64))
		db := c.Get("db").(*gorm.DB)
		adminRepo := repositories.NewAdminRepository(db)
		admin, err := adminRepo.GetAdminByID(userID)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		c.Set("admin", admin)
		return next(c)
	}
}
