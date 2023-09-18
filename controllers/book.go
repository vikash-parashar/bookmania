package controllers

import (
	"myapp/models"

	"gorm.io/gorm"
)

// CreateBook creates a new book in the database.
func CreateBook(db *gorm.DB, book *models.Book) error {
	result := db.Create(&book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetBookByID retrieves a book by its ID.
func GetBookByID(db *gorm.DB, id uint) (*models.Book, error) {
	var book models.Book
	result := db.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

// UpdateBook updates an existing book in the database.
func UpdateBook(db *gorm.DB, book *models.Book) error {
	result := db.Save(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteBook deletes a book by its ID.
func DeleteBook(db *gorm.DB, id uint) error {
	result := db.Delete(&models.Book{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllBooks retrieves all books from the database.
func GetAllBooks(db *gorm.DB) ([]models.Book, error) {
	var books []models.Book
	result := db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}
