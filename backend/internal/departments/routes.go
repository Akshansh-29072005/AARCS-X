package departments

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {

	group := r.Group("/api/v1")

	// Student Creating Route
	group.POST("/departments", h.CreateDepartment)

	// Student Info Getting Route
	group.GET("/departments", h.Read)
}
