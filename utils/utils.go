// utils/utils.go

package utils

import (
	"database/sql"
	"fmt"
	"log"
	"my_bookstore/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// GenerateUUID generates a new UUID.
func GenerateUUID() string {
	return uuid.New().String()
}

// HashPassword generates a bcrypt hash of a password.
func HashPassword(password string) (string, error) {
	// Generate a bcrypt hash of the provided password with a default cost factor
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// MatchWithHashPassword checks if a password matches its bcrypt hash.
func MatchWithHashPassword(providedPassword []byte, storedPassword []byte) bool {
	// Compare the provided password with the stored hash
	err := bcrypt.CompareHashAndPassword(storedPassword, providedPassword)
	if err != nil {
		log.Println(err) // Log the error, but continue processing
		return false
	}
	return true
}

// ExtractEmailFromToken extracts the email from a JWT token.
func ExtractEmailFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Replace "your-secret-key" with your actual JWT secret key
		secretKey := []byte("your-secret-key")
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		return email, nil
	} else {
		return "", fmt.Errorf("Invalid JWT token")
	}
}

// GetUserByID retrieves user data based on their role and renders the appropriate user profile template.
func GetUserByID(c *gin.Context) {
	// Get the user's role from the Gin context
	userRole, _ := c.Get("role")

	// Render different templates based on the user's role
	if userRole == "admin" {
		// Render admin user profile template
		renderAdminUserProfile(c)
	} else if userRole == "user" {
		// Render regular user profile template
		renderRegularUserProfile(c)
	}
}

// renderAdminUserProfile renders the admin user profile template with user data.
func renderAdminUserProfile(c *gin.Context) {
	// Get user data and other necessary information
	userID := c.Param("id") // Assuming you're getting the user ID from the URL

	// Fetch the user's data from your data store
	user, err := getUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user data"})
		return
	}

	// Render the admin user profile template with the user's data
	c.HTML(http.StatusOK, "admin_user_profile.html", gin.H{
		"user": user,
	})
}

// renderRegularUserProfile renders the regular user profile template with user data.
func renderRegularUserProfile(c *gin.Context) {
	// Get user data and other necessary information
	userID := c.Param("id") // Assuming you're getting the user ID from the URL

	// Fetch the user's data from your data store
	user, err := getUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user data"})
		return
	}

	// Render the regular user profile template with the user's data
	c.HTML(http.StatusOK, "regular_user_profile.html", gin.H{
		"user": user,
	})
}

// getUserByID retrieves user data from the database by user ID.
func getUserByID(userID string) (*models.User, error) {
	// Assuming you are using a SQL database (e.g., PostgreSQL, MySQL)
	db, err := sql.Open("your-database-driver", "connection-string")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to the database")
	}
	defer db.Close()

	// Query the database to fetch user data by ID
	query := "SELECT id, email, role FROM users WHERE id = ?"
	var user models.User
	err = db.QueryRow(query, userID).Scan(&user.ID, &user.Email, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("User not found")
		}
		return nil, errors.Wrap(err, "Failed to fetch user data from the database")
	}

	return &user, nil
}
