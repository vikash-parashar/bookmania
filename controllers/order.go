package controllers

import (
	"net/http"

	"my_bookstore/db"
	"my_bookstore/middleware"
	"my_bookstore/models"

	"github.com/gin-gonic/gin"
	// Import other necessary packages
)

// CreateOrder handles the route for creating a new order (admin-only).
func CreateOrder(c *gin.Context) {
	// Check if the user is an admin
	if !middleware.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Parse order data from the request (assuming JSON)
	var newOrder models.Order // Replace with your Order struct

	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the order data (e.g., check if required fields are provided)

	// Implement your validation logic here
	// Example:
	// if newOrder.ProductID == 0 {
	//     c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
	//     return
	// }

	// Save the order to the database
	if err := db.DB.Create(&newOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating order"})
		return
	}

	// Return a success response with the newly created order
	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order": newOrder})
}

// ViewAllOrders retrieves all orders (admin-only).
func ViewAllOrders(c *gin.Context) {
	// Check if the user is an admin
	if !middleware.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Implement logic to retrieve all orders from the database
	var orders []models.Order // Replace with your Order struct

	// Assuming you have a function to retrieve all orders from the database
	if err := db.DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// ViewOrderByUserEmail retrieves orders by user email (admin-only).
func ViewOrderByUserEmail(c *gin.Context) {
	// Check if the user is an admin
	if !middleware.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Extract user email from the request parameters or query parameters
	userEmail := c.Param("email")

	// Implement logic to retrieve orders by user email from the database
	var orders []models.Order // Replace with your Order struct

	// Assuming you have a function to retrieve orders by user email from the database
	if err := db.DB.Where("user_email = ?", userEmail).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving orders by user email"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// UpdateOrderStatus updates the status of an order by ID (admin-only).
func UpdateOrderStatus(c *gin.Context) {
	// Check if the user is an admin
	if !middleware.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Extract order ID from the request parameters
	// orderID := c.Param("id")

	// Implement logic to update the status of the order with the given ID
	// You can get the updated status from the request body or query parameters

	// Example:
	// updatedStatus := c.PostForm("status")

	// Update the order status in the database

	c.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully"})
}

// CancelOrder cancels an order by ID (admin-only).
func CancelOrder(c *gin.Context) {
	// Check if the user is an admin
	if !middleware.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Extract order ID from the request parameters
	// orderID := c.Param("id")

	// Implement logic to cancel the order with the given ID
	// You can update the order status or perform other cancellation actions

	// Example:
	// Update the order status to "Cancelled" in the database

	c.JSON(http.StatusOK, gin.H{"message": "Order cancelled successfully"})
}

// DeleteOrder deletes an order by ID (admin-only).
func DeleteOrder(c *gin.Context) {
	// Check if the user is an admin
	if !middleware.IsAdmin(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Extract order ID from the request parameters
	// orderID := c.Param("id")

	// Implement logic to delete the order with the given ID from the database

	// Example:
	// Delete the order from the database

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
