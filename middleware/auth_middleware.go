package middleware

import (
	"fmt"
	"my_bookstore/models"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("your-secret-key")

// AuthMiddleware is a middleware that handles JWT token validation and user role checking.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract and validate the JWT token from the request
		token, err := extractTokenFromRequest(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		// Verify the token and get its claims
		claims, err := verifyTokenAndGetClaims(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		// Check the user's role from the claims
		userRole := claims["role"].(string)
		if userRole != "admin" && userRole != "user" {
			c.JSON(http.StatusForbidden, gin.H{"message": "Access denied"})
			c.Abort()
			return
		}

		// Set the user's role in the Gin context for later use
		c.Set("role", userRole)

		// Continue to the next middleware or route handler
		c.Next()
	}
}

// GenerateToken generates a new JWT token for the given email address.
func GenerateToken(email string) (string, error) {
	// Set token expiration time to 24 hours from now
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create JWT claims with subject (email) and expiration time
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   email,
	}

	// Create and sign a new JWT token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateToken validates a JWT token and returns the email if valid.
func ValidateToken(tokenString string) (string, error) {
	// Parse and validate the JWT token with the secret key
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// Check for errors and token validity
	if err != nil || !token.Valid {
		return "", err
	}

	// Extract the email (subject) from the token's claims
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("Invalid token claims")
	}

	return claims.Subject, nil
}

// verifyTokenAndGetClaims parses and verifies a JWT token and returns its claims.
func verifyTokenAndGetClaims(tokenString string) (jwt.MapClaims, error) {
	// Parse and verify the JWT token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the token signing method and return the secret key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method")
		}
		return jwtKey, nil
	})

	// Check for errors and token validity
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	// Extract and return the token's claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Invalid token claims")
	}

	return claims, nil
}

// extractTokenFromRequest extracts a JWT token from the request's Authorization header.
func extractTokenFromRequest(req *http.Request) (string, error) {
	// Extract the token from the Authorization header, assuming it's a Bearer token
	tokenHeader := req.Header.Get("Authorization")
	if tokenHeader == "" {
		return "", fmt.Errorf("Token not found")
	}

	// Token format: "Bearer <token>"
	tokenParts := strings.Split(tokenHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return "", fmt.Errorf("Invalid token format")
	}

	return tokenParts[1], nil
}

// IsAdmin checks if the user is an admin.
func IsAdmin(c *gin.Context) bool {
	// Check the user's role or any other criteria to determine if they are an admin
	// For example, you might check if the user has an "admin" role or a specific flag set in their user profile.
	// This implementation depends on your application's user management system.

	// Example: Check if the user has an "admin" role
	user := getCurrentUser(c) // You should implement a function to get the current user
	if user != nil && user.Role == "admin" {
		return true
	}

	return false
}

// getCurrentUser is a placeholder for a function that retrieves the current user from the request.
func getCurrentUser(c *gin.Context) *models.User {
	// Implement logic to retrieve the current user from the request, token, or session
	// Return the user object or nil if the user is not authenticated
	// This implementation depends on your application's authentication system.
	return nil
}
