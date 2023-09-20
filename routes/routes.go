// routes/routes.go

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

	// Use the JWT authentication middleware for specific routes or groups
	// Example:
	// router.Use(middleware.AuthMiddleware())
	// ...
	// Define a route handler for rendering the success page
	router.GET("/success", render.RenderSuccessPage)

	// Define a route handler for rendering the error page
	router.GET("/error", render.RenderErrorPage)
	// Authentication-related routes
	authGroup := router.Group("/auth")

	{
		authGroup.GET("/login", controllers.LoginPage)
		authGroup.POST("/login", controllers.Login)
		authGroup.GET("/register", controllers.RegisterPage)
		authGroup.POST("/register", controllers.Register)
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
