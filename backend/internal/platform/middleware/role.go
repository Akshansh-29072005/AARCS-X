package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type RoleProvider interface {
	GetRolesByID(ctx context.Context, userID int) ([]string, error)
}

func RequireRole(roleProvider RoleProvider, allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := GetLogger(c)
		userID, exists := c.Get(UserIDKey)

		if !exists {
			log.Warn().Msg("UserID not found in context")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		userIDInt, ok := userID.(int)
		if !ok {
			log.Warn().Interface("user_id_value", userID).Msg("UserID type assertion failed")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		if userIDInt == 0 {
			log.Warn().Msg("UserID is 0")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		roles, err := roleProvider.GetRolesByID(c.Request.Context(), userIDInt)
		if err != nil {
			log.Error().Err(err).Int("user_id", userIDInt).Msg("Database error fetching role")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "failed to retrieve user role",
			})
			return
		}

		log.Info().Int("user_id", userIDInt).Str("found_role", strings.Join(roles, ",")).Interface("allowed_roles", allowedRoles).Msg("Checking permissions")

		for _, r := range allowedRoles {
			for _, role := range roles {
				if role == r {
					log.Debug().Str("role", role).Msg("Role authorized")
					c.Next()
					return
				}
			}
		}

		log.Warn().
			Str("actual_role", strings.Join(roles, ",")).
			Interface("allowed_roles", allowedRoles).
			Int("user_id", userIDInt).
			Msg("Insufficient permissions")

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "insufficient permissions",
		})
	}
}
