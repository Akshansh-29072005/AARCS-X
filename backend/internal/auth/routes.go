package auth

import "github.com/gin-gonic/gin"

func RegisteredRoutes(r *gin.Engine, h *Handler) {

	group := r.Group("/api/v1/auth")

	// User Logging In Route
	group.POST("/login")

	// User Logging Out Route
	group.GET("/logout")

	// Testing Route for Auth
	group.GET("/protected")

	// User Info Route
	group.GET("/me")
}
