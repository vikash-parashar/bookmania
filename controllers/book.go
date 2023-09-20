package controllers

import (
	"errors"
	"net/http"

	"my_bookstore/db"
	"my_bookstore/middleware"
	"my_bookstore/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ManageBooks retrieves and manages books (admin-only).
func ManageBooks(c *gin.Context) {
	// Check if the user is an admin
	if !middleware.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Your implementation here
}

// CreateBook creates a new book (admin-only).
func CreateBook(c *gin.Context) {
	// Check if the user is an admin
	if !middleware.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Parse book data from the request (assuming JSON)
	var newBook models.Book // Replace with your Book struct

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the book data (e.g., check if required fields are provided)

	// Implement your validation logic here
	// Example:
	// if newBook.Title == "" {
	//     c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
	//     return
	// }

	// Save the new book to the database
	if err := db.DB.Create(&newBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating book"})
		return
	}

	// Return a success response with the newly created book
	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully", "book": newBook})
}

// GetBookByID retrieves book information by ID (admin-only).
func GetBookByID(c *gin.Context) {
	// Check if the user is an admin
	if !middleware.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Extract book ID from the request parameters
	bookID := c.Param("id")

	// Implement logic to retrieve book information by ID from the database

	var book models.Book // Replace with your Book struct

	if err := db.DB.Where("id = ?", bookID).First(&book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook updates book information by ID (admin-only).
func UpdateBook(c *gin.Context) {
	// Check if the user is an admin
	if !middleware.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Extract book ID from the request parameters
	// bookID := c.Param("id")

	// Parse updated book data from the request (assuming JSON)
	var updatedBook models.Book // Replace with your Book struct

	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the updated book data

	// Implement your validation logic here

	// Update the book information in the database

	// Example:
	// if err := db.DB.Model(&Book{}).Where("id = ?", bookID).Updates(updatedBook).Error; err != nil {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating book"})
	//     return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

// DeleteBook deletes a book by ID (admin-only).
func DeleteBook(c *gin.Context) {
	// Check if the user is an admin
	if !middleware.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Extract book ID from the request parameters
	// bookID := c.Param("id")

	// Implement logic to delete the book with the given ID from the database

	// Example:
	// if err := db.DB.Delete(&Book{}, bookID).Error; err != nil {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting book"})
	//     return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
