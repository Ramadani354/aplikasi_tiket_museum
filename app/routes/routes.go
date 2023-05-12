package routes

import (
	"github.com/Ramadani354/tiket_museum/app/controllers"
	"github.com/Ramadani354/tiket_museum/app/middlewares"
	"github.com/Ramadani354/tiket_museum/app/repositories"
	"github.com/Ramadani354/tiket_museum/app/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	printRepo := repositories.NewPrintRepository(db)
	printService := services.NewPrintService(printRepo)
	printController := controllers.NewPrintController(printService)

	adminController := controllers.NewAdminController(db)
	ticketController := controllers.NewTicketController(db)

	authGroup := e.Group("/auth")
	authGroup.POST("/register", adminController.Register)
	authGroup.POST("/login", adminController.Login)

	authenticated := e.Group("/api")
	authenticated.Use(middlewares.JWTMiddleware())
	authenticated.POST("/tickets", ticketController.CreateTicket)
	authenticated.GET("/tickets", ticketController.GetAllTickets)
	authenticated.GET("/tickets/:id", ticketController.GetTicket)
	authenticated.PUT("/tickets/:id", ticketController.UpdateTicket)
	authenticated.DELETE("/tickets/:id", ticketController.DeleteTicket)

	adminGroup := e.Group("/admin")
	adminGroup.Use(middlewares.JWTMiddleware())
	adminGroup.GET("/tickets/kuota", ticketController.GetTicketQuota)
	adminGroup.POST("/print", printController.CreatePrint)
}
