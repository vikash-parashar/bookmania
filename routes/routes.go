package routes

import (
	"my_bookstore/controllers"
	"my_bookstore/middleware"
	"my_bookstore/render"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Use the error handling middleware globally
	router.Use(middleware.ErrorMiddleware())

	// Define a route handler for rendering the success page
	router.GET("/success", render.RenderSuccessPage)

	// Define a route handler for rendering the error page
	router.GET("/error", render.RenderErrorPage)

	// Authentication-related routes
	authGroup := router.Group("/auth")
	{
		// Render the login page
		authGroup.GET("/login", controllers.LoginPage)

		// Handle user login
		authGroup.POST("/login", controllers.Login)

		// Render the registration page
		authGroup.GET("/register", controllers.RegisterPage)

		// Handle user registration
		authGroup.POST("/register", controllers.Register)

		// Handle user logout
		authGroup.POST("/logout", controllers.Logout)
	}

	// User-related routes
	userGroup := router.Group("/users")
	{
		// Use the AuthMiddleware for user-related routes
		userGroup.Use(middleware.AuthMiddleware())

		// Create a new user
		userGroup.POST("/", controllers.CreateUser)

		// Retrieve user information by ID
		userGroup.GET("/:id", controllers.GetUserByID)

		// Update user information by ID
		userGroup.PUT("/:id", controllers.UpdateUser)

		// Delete user by ID
		userGroup.DELETE("/:id", controllers.DeleteUser)

		// View order history for a user
		userGroup.GET("/:id/orders", controllers.ViewOrderHistory)

		// Buy a book for a user
		userGroup.POST("/:id/orders/buy", controllers.BuyBook)

		// Rent a book for a user
		userGroup.POST("/:id/orders/rent", controllers.RentBook)
	}

	// Admin-related routes
	adminGroup := router.Group("/admin")
	{
		// Use the AuthMiddleware for admin-related routes (if needed)
		// adminGroup.Use(middleware.AuthMiddleware())

		// View and manage books for admins
		adminGroup.GET("/books", controllers.ManageBooks)
		adminGroup.POST("/books", controllers.CreateBook)
		adminGroup.GET("/books/:id", controllers.GetBookByID)
		adminGroup.PUT("/books/:id", controllers.UpdateBook)
		adminGroup.DELETE("/books/:id", controllers.DeleteBook)

		// View and manage orders for admins
		adminGroup.GET("/orders", controllers.ViewAllOrders)
		adminGroup.GET("/orders/user/:email", controllers.ViewOrderByUserEmail)
		adminGroup.PUT("/orders/:id", controllers.UpdateOrderStatus)
		adminGroup.POST("/orders/cancel/:id", controllers.CancelOrder)
		adminGroup.DELETE("/orders/:id", controllers.DeleteOrder)

		// Generate a sales report for admins
		adminGroup.GET("/sales", controllers.GenerateSalesReport)

		// View and manage users for admins
		adminGroup.GET("/users", controllers.GetAllUsers)
		adminGroup.POST("/users", controllers.CreateUser)
		adminGroup.GET("/users/:id", controllers.GetUserByID)
		adminGroup.PUT("/users/:id", controllers.UpdateUser)
		adminGroup.DELETE("/users/:id", controllers.DeleteUser)
	}

	// Add other routes as needed
}
