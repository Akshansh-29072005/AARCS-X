package teachers

import (
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/middleware"
	"github.com/gin-gonic/gin"
)

func RegisteredRoutes(r *gin.Engine, h *Handler) {

	group := r.Group("/api/v1")
	group.Use(middleware.AuthMiddleware())
	group.Use(middleware.RequireRole("institution"))

	// Teacher Creating Route
	group.POST("/teachers", h.CreateTeacher)

	// Teacher Info Getting Route
	group.GET("/teachers", h.Read)
}
