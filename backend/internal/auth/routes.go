package auth

import (
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/middleware"
	"github.com/gin-gonic/gin"
)

func RegisteredRoutes(r *gin.Engine, h *Handler) {

	api := r.Group("/api/v1/auth")

	// User Registering Route
	api.POST("/register", h.RegisterInstitution)

	// User Logging In Route
	api.POST("/login", h.Login)

	protected := api.Group("/protected")
	protected.Use(middleware.AuthMiddleware())

	// User Info Route
	protected.GET("/me", h.Me)

	// User Logging Out Route
	//protected.POST("/logout", h.Logout)
}
