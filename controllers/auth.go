// controllers/auth_controller.go

package controllers

import (
	"log"
	"my_bookstore/config"
	"my_bookstore/db"
	"my_bookstore/models"
	"my_bookstore/render"
	"my_bookstore/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Login handles user login and JWT token generation.
func Login(c *gin.Context) {
	if err := render.RenderTemplate(c.Writer, "login", nil); err != nil {
		log.Println(err)
	}

	var loginData struct {
		Email    string `json:"email"`
		Password []byte `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Query the database to find the user by email
	var user models.User
	if err := db.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare the provided password with the hashed password in the database
	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
	// 	return
	// }
	utils.CheckPasswordHash(string(loginData.Password), string(user.Password))

	// Create JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign the token with the JWT secret
	tokenString, err := token.SignedString([]byte(config.AppConfig.JWTSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}

	// Create and send a refresh token (if needed)
	// ...

	// Respond with the JWT token
	c.JSON(http.StatusOK, gin.H{"token": tokenString, "message": "Logged in successfully"})
}

// Login handles user login and JWT token generation.
func LoginPage(c *gin.Context) {
	if err := render.RenderTemplate(c.Writer, "login", nil); err != nil {
		log.Println(err)
	}
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
