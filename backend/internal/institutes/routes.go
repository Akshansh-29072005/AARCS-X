package institutes

import (
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *Handler) {

	group := r.Group("/api/v1")
	group.Use(middleware.AuthMiddleware())
	group.Use(middleware.RequireRole("institution"))

	// Institution Creating Route
	group.POST("/institutions", h.CreateInstitute)

	// Institution Info Getting Route
	group.GET("/institutions", h.Read)

	// Institution Info Getting by ID Route
	group.GET("/institutions/:id", h.ReadByID)
}
