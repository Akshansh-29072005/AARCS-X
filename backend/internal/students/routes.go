package students

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {

	group := r.Group("/api/v1")

	// Student Creating Route
	group.POST("/students", h.CreateStudent)

	// Student Info Getting Route
	group.GET("/students", h.Read)
}
