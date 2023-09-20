package controllers

import (
	"errors"
	"net/http"

	"my_bookstore/db"
	"my_bookstore/middleware"
	"my_bookstore/models"
	"my_bookstore/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ViewOrderHistory retrieves the order history of a user by ID.
func ViewOrderHistory(c *gin.Context) {
	// Your implementation here
}

// BuyBook allows a user to make a purchase.
func BuyBook(c *gin.Context) {
	// Your implementation here
}

// RentBook allows a user to rent a book.
func RentBook(c *gin.Context) {
	// Your implementation here
}

// GetAllUsers retrieves all users (admin-only).
func GetAllUsers(c *gin.Context) {
	// Check if the user is an admin
	if !middleware.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Your logic to retrieve all users from the database
	var users []models.User // Replace with your User struct

	// Assuming you have a function to retrieve all users from the database
	if err := db.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// CreateUser handles the creation of a new user (admin route).
func CreateUser(c *gin.Context) {
	// Check if the user is an admin
	if !middleware.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Parse user data from the request (assuming JSON)
	var newUser models.User // You should define the User struct according to your application's needs
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if a user with the given email already exists in the database
	var existingUser models.User
	if err := db.DB.Where("email = ?", newUser.Email).First(&existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User with the email does not exist, proceed with registration

			// Generate hashed password
			hashedPassword, err := utils.HashPassword(string(newUser.Password))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
				return
			}

			// Set the hashed password
			newUser.Password = []byte(hashedPassword)

			// Create the new user
			if err := db.DB.Create(&newUser).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
				return
			}

			c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": newUser})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "User with this email already exists"})
}

// GetUserByID retrieves user information by ID (admin route).
func GetUserByID(c *gin.Context) {
	// Check if the user is an admin
	if !middleware.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Retrieve the user ID from the URL params
	userID := c.Param("id")

	// Your implementation here, e.g., fetch user by userID

	// Return user data in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID",
		"id":      userID,
		// Add user data here
	})
}

// UpdateUser updates user information by ID (admin route).
func UpdateUser(c *gin.Context) {
	// Check if the user is an admin
	if !middleware.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Parse user data from the request (assuming JSON)
	var updatedUser models.User // Replace with your User struct

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user with the given ID exists
	var existingUser models.User // Replace with your User struct

	if err := db.DB.Where("id = ?", updatedUser.ID).First(&existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Update user information
	existingUser.FirstName = updatedUser.FirstName
	existingUser.LastName = updatedUser.LastName
	existingUser.Phone = updatedUser.Phone
	existingUser.Email = updatedUser.Email

	// Save the updated user to the database
	if err := db.DB.Save(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": existingUser})
}

// DeleteUser deletes a user by ID (admin route).
func DeleteUser(c *gin.Context) {
	// Check if the user is an admin
	if !middleware.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Get the user ID from the request parameters
	userID := c.Param("id")

	// Check if the user with the given ID exists
	var existingUser models.User // Replace with your User struct

	if err := db.DB.Where("id = ?", userID).First(&existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Delete the user from the database
	if err := db.DB.Delete(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
