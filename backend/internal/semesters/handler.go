package semesters

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

func (h *Handler) CreateSemester(c *gin.Context) {

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "semesters_handler").
		Msg("Received request to create semester")

	var req CreateSemesterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	semester, err := h.service.CreateSemesters(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, semester)

	log.Info().
		Str("component", "semesters_handler").
		Msg("Semester created successfully")
}

func (h *Handler) Read(c *gin.Context) {

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "semesters_handler").
		Msg("Received request to read semesters")

	var req GetSemestersRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.service.GetSemesters(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)

	log.Info().
		Str("component", "semesters_handler").
		Msg("Semesters retrieved successfully")
}
