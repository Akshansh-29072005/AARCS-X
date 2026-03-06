package teachers

import (
	"net/http"

	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateTeacher(c *gin.Context) {

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "teachers_handler").
		Msg("Received request to create teacher")

	var req CreateTeacherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teacher, err := h.service.CreateTeacher(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, teacher)

	log.Info().
		Str("component", "teachers_handler").
		Msg("Teacher created successfully")
}

func (h *Handler) Read(c *gin.Context) {

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "teachers_handler").
		Msg("Received request to read teachers")

	var req GetTeachersRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.service.GetTeachers(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)

	log.Info().
		Str("component", "teachers_handler").
		Msg("Teachers retrieved successfully")
}
