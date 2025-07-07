package db

import (
	"balzac/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("balzac.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	DB = db

	// Auto-migrate the User model
	DB.AutoMigrate(&models.User{})
}
