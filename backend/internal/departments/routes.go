package departments

import (
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *Handler, rp middleware.RoleProvider) {

	group := r.Group("/api/v1")
	group.Use(middleware.AuthMiddleware())
	group.Use(middleware.RequireRole(rp, "institution"))

	// Student Creating Route
	group.POST("/departments", h.CreateDepartment)

	// Student Info Getting Route
	group.GET("/departments", h.Read)
}
