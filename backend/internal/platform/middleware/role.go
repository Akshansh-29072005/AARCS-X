package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const RoleInstitutionAdmin = "institution_admin"

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString(RoleInstitutionAdmin)

		for _, r := range allowedRoles {
			if role == r {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "insufficient permissions",
		})
	}
}
