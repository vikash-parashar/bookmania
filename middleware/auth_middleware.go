// middleware/auth_middleware.go

package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Your JWT authentication logic here
		// Verify the JWT token in the request headers
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Verify the JWT token using your JWT library of choice
		// If the token is valid, set user information in the context
		// If the token is invalid, return an unauthorized response
		// Example:
		// user, err := jwt.VerifyToken(token)
		// if err != nil {
		//     c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		//     c.Abort()
		//     return
		// }
		// c.Set("user", user)

		c.Next()
	}
}

var jwtKey = []byte("your-secret-key")

// GenerateToken generates a new JWT token for the given email address.
func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token and returns the email if valid.
func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return "", err
	}

	return claims.Subject, nil
}
