package subjects

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {

	group := r.Group("/api/v1")

	// Student Creating Route
	group.POST("/subjects", h.CreateSubject)

	// Student Info Getting Route
	group.GET("/subjects", h.Read)
}
