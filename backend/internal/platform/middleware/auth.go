package middleware

import (
	"net/http"
	"strings"

	"github.com/Akshansh-29072005/AARCS-X/backend/internal/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header missing",
			})
			return
		}

		// Check Bearer format
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header format must be Bearer {token}",
			})
			return
		}

		tokenString := parts[1]

		// Parse & validate token
		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			return
		}

		// Store identity in context
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
		c.Set("ref_id", claims.RefID)

		// Continue request
		c.Next()
	}
}
