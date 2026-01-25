package institutes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {

	group := r.Group("/api/v1")

	// Student Creating Route
	group.POST("/institutions", h.CreateInstitute)

	// Student Info Getting Route
	group.GET("/institutions", h.Read)
}
