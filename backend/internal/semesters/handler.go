package semesters

import (
	"fmt"
	"net/http"
	"strconv"

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

func (h *Handler) ReadByID(c *gin.Context) {

	idStr := c.Param("id")

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "semesters_handler").
		Str("id", idStr).
		Msg("Received request to read semester by ID")

	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id parameter is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id parameter must be a valid integer"})
		return
	}

	response, cacheHit, err := h.service.GetSemesterByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	if cacheHit {
		log.Info().
			Str("cache_status", "hit").
			Str("cache_key", fmt.Sprintf("semester:v1:%d", id)).
			Msg("cache hit")
	} else {
		log.Info().
			Str("cache_status", "miss").
			Str("cache_key", fmt.Sprintf("semester:v1:%d", id)).
			Msg("cache miss")
	}

	c.JSON(http.StatusOK, response)

	log.Info().
		Str("component", "semesters_handler").
		Str("id", idStr).
		Msg("Semester retrieved successfully")
}