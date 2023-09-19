// middleware/auth_middleware.go

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
