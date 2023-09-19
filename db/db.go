package db

import (
	"my_bookstore/config" // Import your configuration package
	"my_bookstore/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	dbURL := config.AppConfig.GetDatabaseURL()

	// Open a database connection
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	DB.AutoMigrate(&models.Address{}, &models.User{}, &models.Author{}, &models.Book{}, &models.Contact{}, &models.Order{})

}
