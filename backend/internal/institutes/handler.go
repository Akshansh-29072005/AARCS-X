package institutes

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

func (h *Handler) CreateInstitute(c *gin.Context) {

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "institutes_handler").
		Msg("Received request to create institute")

	var req CreateInstitutionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	institute, err := h.service.CreateInstitution(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, institute)

	log.Info().
		Str("component", "institutes_handler").
		Msg("Institute created successfully")
}

func (h *Handler) Read(c *gin.Context) {

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "institutes_handler").
		Msg("Received request to read institutes")

	var req GetInstitutionRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.service.GetInstitutions(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)

	log.Info().
		Str("component", "institutes_handler").
		Msg("Institutes retrieved successfully")
}
