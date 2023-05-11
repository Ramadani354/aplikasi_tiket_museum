package routes

import (
	"github.com/Ramadani354/tiket_museum/app/controllers"
	"github.com/Ramadani354/tiket_museum/app/middlewares"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {

	// Inisialisasi controller dengan meneruskan objek db
	adminController := controllers.NewAdminController(db)
	ticketController := controllers.NewTicketController(db)
	printController := controllers.NewPrintController(db)

	// Grup rute untuk auth
	authGroup := e.Group("/auth")
	authGroup.POST("/register", adminController.Register)
	authGroup.POST("/login", adminController.Login)

	// Grup rute yang membutuhkan otentikasi
	authenticated := e.Group("/api")
	authenticated.Use(middlewares.JWTMiddleware()) // Middleware otentikasi JWT
	authenticated.POST("/tickets", ticketController.CreateTicket)
	authenticated.GET("/tickets", ticketController.GetAllTickets)
	authenticated.GET("/tickets/:id", ticketController.GetTicket)
	authenticated.PUT("/tickets/:id", ticketController.UpdateTicket)
	authenticated.DELETE("/tickets/:id", ticketController.DeleteTicket)

	// Grup rute admin
	adminGroup := e.Group("/admin")
	adminGroup.Use(middlewares.JWTMiddleware()) // Middleware otentikasi JWT untuk admin
	adminGroup.GET("/tickets/kuota", ticketController.GetTicketQuota)
	adminGroup.POST("/print", printController.CreatePrint)

}
