package middlewares

import (
	"net/http"
	"strings"

	"github.com/Ramadani354/tiket_museum/utils"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
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

			c.Set("user_id", claims["user_id"])
			return next(c)
		}
	}
}
