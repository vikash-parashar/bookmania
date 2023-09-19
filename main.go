package main

import (
	"my_bookstore/config"
	"my_bookstore/db"
	"my_bookstore/middleware"
	"my_bookstore/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	// Load the application configuration
	config.LoadConfig()

	// Initialize the database connection
	db.Init()
}

func main() {

	// Initialize Gin router
	router := gin.Default()

	// Use the error handling middleware globally
	router.Use(middleware.ErrorMiddleware())

	// Initialize routes
	routes.SetupRoutes(router)
	// Start the server
	router.Run(":8080")

}
