// controllers/auth_controller.go

package controllers

import (
	"errors"
	"log"
	"my_bookstore/db"

	// "my_bookstore/middleware"
	"my_bookstore/models"
	"my_bookstore/render"
	"my_bookstore/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LoginPage handles rendering the login page.
func LoginPage(c *gin.Context) {
	if err := render.RenderTemplate(c.Writer, "login", nil); err != nil {
		log.Printf("Error rendering login page: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
}

// Login handles user login and JWT token generation.
func Login(c *gin.Context) {
	var loginData struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}

	if err := c.ShouldBind(&loginData); err != nil {
		log.Printf("Error binding form data: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Now you can access loginData.Email and loginData.Password as form data

	// Query the database to find the user by email
	var storedUser models.User
	if err := db.DB.Where("email = ?", loginData.Email).First(&storedUser).Error; err != nil {
		log.Printf("Error querying user: %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare the provided password with the hashed password in the database
	if utils.MatchWithHashPassword([]byte(loginData.Password), storedUser.Password) {
		// Generate a JWT token using jwtutil.GenerateToken
		// tokenString, err := middleware.GenerateToken(loginData.Email)
		// if err != nil {
		// 	log.Printf("Error generating JWT token: %v\n", err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating JWT token"})
		// 	return
		// }

		// Send the token as a response

		// c.JSON(http.StatusOK, gin.H{"token": tokenString})

		c.HTML(http.StatusOK, "success.log.html", gin.H{"message": "you are now logged in"})
	} else {
		// Send a success message and render the success page
		c.HTML(http.StatusOK, "error.log.html", gin.H{"message": "you have entered invalid email id or password"})
		log.Println("Invalid credentials")
	}
}

// RegisterPage handles rendering the register page.
func RegisterPage(c *gin.Context) {
	if err := render.RenderTemplate(c.Writer, "register", nil); err != nil {
		log.Printf("Error rendering register page: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
}

// Register handles user registration and JWT token generation.
func Register(c *gin.Context) {
	var loginData struct {
		FirstName       string `form:"firstName"`
		LastName        string `form:"lastName"`
		Phone           string `form:"phone"`
		Email           string `form:"email"`
		Password        string `form:"password"`
		ConfirmPassword string `form:"confirmPassword"`
	}

	// Bind the form data to the loginData struct
	if err := c.ShouldBind(&loginData); err != nil {
		log.Printf("Error binding form data: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Check if the password and confirmPassword match
	if loginData.Password != loginData.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	// Check if a user with the given email already exists in the database
	var existingUser models.User
	if err := db.DB.Where("email = ?", loginData.Email).First(&existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User with the email does not exist, proceed with registration

			// Generate hashed password
			hashedPassword, err := utils.HashPassword(loginData.Password)
			if err != nil {

				// For error cases, render the error page
				c.HTML(http.StatusOK, "error.reg.html", nil)
				log.Println("error : error hashing password")
				return
			}

			// Create a new user model
			newUser := models.User{
				FirstName: loginData.FirstName,
				LastName:  loginData.LastName,
				Phone:     loginData.Phone,
				Email:     loginData.Email,
				Password:  []byte(hashedPassword), // Store the hashed password
			}

			// Store the user in the database
			if err := db.DB.Create(&newUser).Error; err != nil {

				// For error cases, render the error page
				c.HTML(http.StatusOK, "error.reg.html", gin.H{"error": "error while creating user"})
				log.Println("error : error while creating user")
				return
			}

			// Send a success message and render the success page
			c.HTML(http.StatusOK, "success.reg.html", gin.H{"message": "User created successfully"})

			// Send a success message in the response
			log.Println("message : User created successfully")
			return // Return to prevent further execution
		}

		// For error cases, render the error page
		c.HTML(http.StatusOK, "error.reg.html", nil)
		// Other error occurred while querying the database
		log.Printf("error querying user: %v\n", err)
		log.Println("error : database error")
		return
	}

	// For error cases, render the error page
	c.HTML(http.StatusOK, "error.reg.html", gin.H{"error": "user with this email is already exists"})
	log.Println("error : user with this email is already exists")
}

// Logout invalidates the JWT token (client-side implementation).
func Logout(c *gin.Context) {
	// In a stateless JWT-based authentication system, logout is typically handled client-side.
	// You can't invalidate the JWT token on the server-side without additional complexity.

	// To perform a logout operation, the client can simply discard the JWT token.
	// The JWT token will automatically become invalid once it expires.

	// Optionally, you can implement token revocation using a token blacklist on the server-side.
	// However, this adds complexity to the authentication system.

	// Send a response indicating successful logout (optional).
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
