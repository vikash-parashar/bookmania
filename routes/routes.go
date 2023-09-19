// routes/routes.go

package routes

import (
	"my_bookstore/controllers"
	"my_bookstore/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Use the error handling middleware globally
	router.Use(middleware.ErrorMiddleware())

	// Use the JWT authentication middleware for specific routes or groups
	// Example:
	// router.Use(middleware.AuthMiddleware())
	// ...

	// Authentication-related routes
	authGroup := router.Group("/auth")
	{
		authGroup.GET("/login", controllers.LoginPage)
		authGroup.POST("/login", controllers.Login)
		authGroup.POST("/logout", controllers.Logout)
	}

	// User-related routes
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", controllers.CreateUser)
		userGroup.GET("/:id", controllers.GetUserByID)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}

	// Add other routes as needed
}
