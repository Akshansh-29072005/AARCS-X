package middleware

import (
	"net/http"
	"strings"

	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := GetLogger(c)

		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Warn().Msg("Authorization header missing")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header missing",
			})
			return
		}

		// Check Bearer format
		if !strings.HasPrefix(authHeader, "Bearer ") {
			log.Warn().Msg("Invalid authorization header format")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header format must be Bearer {token}",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse & validate token
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			log.Warn().Err(err).Msg("Invalid or expired token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			return
		}

		// Store identity in context
		c.Set(UserIDKey, claims.UserID)

		// Continue request
		c.Next()
	}
}
