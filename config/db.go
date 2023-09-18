package config

import (
	"log"
	"myapp/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Exported database connection pool
func SetupDB() {
	dsn := "host=localhost user=postgres password=vikash dbname=test sslmode=disable"

	// Open a database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
	// Automatically create tables based on your struct models
	// For example, if you have a User struct, GORM will create a "users" table in the database
	err = db.AutoMigrate(&models.User{}, &models.Book{}, &models.Order{})
	if err != nil {
		log.Fatalln(err)
	}
}
