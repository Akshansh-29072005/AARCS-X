package students

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
	r.POST("/api/students", h.Create)
}
