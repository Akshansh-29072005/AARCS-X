package teachers

import "github.com/gin-gonic/gin"

func RegisteredRoutes(r *gin.Engine, h *Handler) {
	group := r.Group("/api/v1")

	// Teacher Creating Route
	group.POST("/teachers", h.CreateTeacher)
}
