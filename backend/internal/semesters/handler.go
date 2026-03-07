package semesters

import (
	"net/http"

	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/errors"
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
		c.Error(errors.BadRequest("invalid request body", err))
		return
	}

	semester, err := h.service.CreateSemesters(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
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
		c.Error(errors.BadRequest("invalid request query", err))
		return
	}

	response, err := h.service.GetSemesters(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)

	log.Info().
		Str("component", "semesters_handler").
		Msg("Semesters retrieved successfully")
}
