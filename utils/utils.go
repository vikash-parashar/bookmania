// utils/utils.go

package utils

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// GenerateUUID generates a new UUID.
func GenerateUUID() string {
	return uuid.New().String()
}

// HashPassword generates a bcrypt hash of a password.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash checks if a password matches its bcrypt hash.
func MatchWithHashPassword(providedPassword []byte, storedPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(storedPassword, providedPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func ExtractEmailFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// You should replace "your-secret-key" with your actual JWT secret key
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
