package subjects

import (
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *Handler, rp middleware.RoleProvider) {

	group := r.Group("/api/v1")
	group.Use(middleware.AuthMiddleware())
	group.Use(middleware.RequireRole(rp, "institution"))

	// Subject Creating Route
	group.POST("/subjects", h.CreateSubject)

	// Subject Info Getting Route
	group.GET("/subjects", h.Read)

	// Subject Info Getting Route by ID
	group.GET("/subjects/:id", h.ReadByID)
}
