package departments

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

func (h *Handler) CreateDepartment(c *gin.Context) {

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "departments_handler").
		Msg("Received request to create department")
	
	var req CreateDepartmentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	department, err := h.service.CreateDepartment(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, department)

	log.Info().
		Str("component", "departments_handler").
		Msg("Department created successfully")
}

func (h *Handler) Read(c *gin.Context) {	

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "departments_handler").
		Msg("Received request to read departments")

	var req GetDepartmentRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.service.GetDepartments(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
	log.Info().
		Str("component", "departments_handler").
		Msg("Departments retrieved successfully")
}
