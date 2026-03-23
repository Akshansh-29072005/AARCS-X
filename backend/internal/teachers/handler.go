package teachers

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

func (h *Handler) CreateTeacher(c *gin.Context) {

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "teachers_handler").
		Msg("Received request to create teacher")

	var req CreateTeacherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.BadRequest("invalid request body", err))
		return
	}

	teacher, err := h.service.CreateTeacher(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
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
		c.Error(errors.BadRequest("invalid query parameters", err))
		return
	}

	response, err := h.service.GetTeachers(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)

	log.Info().
		Str("component", "teachers_handler").
		Msg("Teachers retrieved successfully")
}

func (h *Handler) ReadByID(c *gin.Context) {

	idStr := c.Param("id")

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "teachers_handler").
		Str("id", idStr).
		Msg("Received request to read teacher by ID")

	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id parameter is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id parameter must be a valid integer"})
		return
	}

	response, cacheHit, err := h.service.GetTeacherByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	if cacheHit {
		log.Info().
			Str("cache_status", "hit").
			Str("cache_key", fmt.Sprintf("teacher:v1:%d", id)).
			Msg("cache hit")
	} else {
		log.Info().
			Str("cache_status", "miss").
			Str("cache_key", fmt.Sprintf("teacher:v1:%d", id)).
			Msg("cache miss")
	}

	c.JSON(http.StatusOK, response)

	log.Info().
		Str("component", "teachers_handler").
		Str("id", idStr).
		Msg("Teacher retrieved successfully")
}