package semesters

import (
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *Handler) {

	group := r.Group("/api/v1")
	group.Use(middleware.AuthMiddleware())
	group.Use(middleware.RequireRole("institution"))

	// Student Creating Route
	group.POST("/semesters", h.CreateSemester)

	// Student Info Getting Route
	group.GET("/semesters", h.Read)
}
