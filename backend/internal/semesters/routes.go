package semesters

import (
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *Handler, rp middleware.RoleProvider) {

	group := r.Group("/api/v1")
	group.Use(middleware.AuthMiddleware())
	group.Use(middleware.RequireRole(rp, "institution"))

	// Semesters Creating Route
	group.POST("/semesters", h.CreateSemester)

	// Semesters Info Getting Route
	group.GET("/semesters", h.Read)

	// Semesters Info Getting Route by ID
	group.GET("/semesters/:id", h.ReadByID)
}
