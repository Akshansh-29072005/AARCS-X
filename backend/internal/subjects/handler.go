package subjects

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

func (h *Handler) CreateSubject(c *gin.Context) {

	log := middleware.GetLogger(c)

	log.Info().
		Str("component", "subjects_handler").
		Msg("Received request to create subject")
	
	var req CreateSubjectRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subject, err := h.service.CreateSubject(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, subject)

	log.Info().
		Str("component", "subjects_handler").
		Msg("Subject created successfully")
}

func (h *Handler) Read(c *gin.Context) {
	var req GetSubjectRequest

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "subjects_handler").
		Msg("Received request to read subjects")

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.service.GetSubjects(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)

	log.Info().
		Str("component", "subjects_handler").
		Msg("Subjects retrieved successfully")
}
