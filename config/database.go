package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/Ramadani354/tiket_museum/app/models"
)

var (
	db *gorm.DB
)

func InitDB() error {
	database, err := gorm.Open(mysql.Open("root:zevo313354@tcp(localhost:3306)/tiket_museum?parseTime=true"), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed", err)
		return err
	} else {
		log.Println("Connection Established")
	}

	db = database

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func AutoMigrate() error {
	err := db.AutoMigrate(
		&models.Admin{},
	)
	if err != nil {
		return err
	}

	return nil
}
