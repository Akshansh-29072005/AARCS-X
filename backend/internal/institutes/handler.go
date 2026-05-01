package institutes

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

func (h *Handler) CreateInstitute(c *gin.Context) {

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "institutes_handler").
		Msg("Received request to create institute")

	var req CreateInstitutionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.BadRequest("invalid request body", err))
		return
	}

	institute, err := h.service.CreateInstitute(c.Request.Context(), req, c.GetInt("user_id"))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, institute)

	log.Info().
		Str("component", "institutes_handler").
		Msg("Institute created successfully")
}

func (h *Handler) MakeAdmin(c *gin.Context) {

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "institutes_handler").
		Msg("Received request to make admin for an institute")

	var req MakeAdminRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.BadRequest("invalid request body", err))
		return
	}

	admin, err := h.service.MakeAdmin(c.Request.Context(), req, c.GetInt(middleware.UserIDKey))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, admin)

	log.Info().
		Str("component", "institutes_handler").
		Msg("Admin created successfully")
}

func (h *Handler) Read(c *gin.Context) {

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "institutes_handler").
		Msg("Received request to read institutes")

	var req GetInstitutionRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.Error(errors.BadRequest("invalid query parameters", err))
		return
	}

	response, err := h.service.GetInstitutions(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)

	log.Info().
		Str("component", "institutes_handler").
		Msg("Institutes retrieved successfully")
}

func (h *Handler) ReadByID(c *gin.Context) {

	idStr := c.Param("id")

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "institutes_handler").
		Str("id", idStr).
		Msg("Received request to read institute by ID")

	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id parameter is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id parameter must be a valid integer"})
		return
	}

	response, cacheHit, err := h.service.GetInstitutionByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	if cacheHit {
		log.Info().
			Str("cache_status", "hit").
			Str("cache_key", fmt.Sprintf("institution:v1:%d", id)).
			Msg("cache hit")
	} else {
		log.Info().
			Str("cache_status", "miss").
			Str("cache_key", fmt.Sprintf("institution:v1:%d", id)).
			Msg("cache miss")
	}

	c.JSON(http.StatusOK, response)

	log.Info().
		Str("component", "institutes_handler").
		Str("id", idStr).
		Msg("Institute retrieved successfully")
}