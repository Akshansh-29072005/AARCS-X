package semesters

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {

	group := r.Group("/api/v1")

	// Student Creating Route
	group.POST("/semesters", h.CreateSemester)

	// Student Info Getting Route
	group.GET("/semesters", h.Read)
}
