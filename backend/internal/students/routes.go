package students

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {

	// Student Creating Route
	r.POST("/api/students", h.Create)

	// Student Getting Route
	r.GET("api/students", h.Read)
}
