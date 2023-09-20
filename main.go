package main

import (
	"my_bookstore/config"
	"my_bookstore/db"
	"my_bookstore/middleware"
	"my_bookstore/routes"
	"net/http"

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

	// Serve static files (CSS, fonts, JS, images)
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*") // Adjust the path to your HTML template files

	// Define a route handler for rendering the success page
	router.GET("/success", func(c *gin.Context) {
		c.HTML(http.StatusOK, "success.reg.html", nil)

		c.HTML(http.StatusOK, "success.log.html", nil)
	})

	// Define a route handler for rendering the error page
	router.GET("/error", func(c *gin.Context) {
		c.HTML(http.StatusOK, "error.reg.html", nil)
		c.HTML(http.StatusOK, "error.log.html", nil)
	})
	// Initialize routes
	routes.SetupRoutes(router)
	// Start the server
	router.Run(":8080")

}
