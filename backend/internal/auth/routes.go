package auth

import "github.com/gin-gonic/gin"

// RegisterRoutes registers auth HTTP routes.
func RegisterRoutes(r *gin.Engine, h *Handler) {
	group := r.Group("/api/v1/auth")
	{
		group.POST("/login", h.Login)
		group.POST("/refresh", h.Refresh)
		group.POST("/logout", h.Logout)
	}
}
