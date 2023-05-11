package main

import (
	"log"

	"github.com/Ramadani354/tiket_museum/app/models"
	"github.com/Ramadani354/tiket_museum/app/routes"
	"github.com/Ramadani354/tiket_museum/config"
	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize instance of Echo
	e := echo.New()

	// Initialize the database
	err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Get the database instance
	db := config.GetDB()
	db.AutoMigrate(&models.Admin{}, &models.Login{}, &models.Print{}, &models.Ticket{})

	// Register routes with Echo and pass the database instance
	routes.RegisterRoutes(e, db)

	// Start the server
	e.Start(":8080")
}
