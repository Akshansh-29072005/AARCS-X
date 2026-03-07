package students

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

func (h *Handler) CreateStudent(c *gin.Context) {

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "students_handler").
		Msg("Received request to create student")

	var req CreateStudentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.BadRequest("invalid request body", err))
		return
	}

	student, err := h.service.CreateStudent(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, student)

	log.Info().
		Str("component", "students_handler").
		Msg("Student created successfully")
}

func (h *Handler) Read(c *gin.Context) {

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "students_handler").
		Msg("Received request to read students")

	var req GetStudentsRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.Error(errors.BadRequest("invalid query parameters", err))
		return
	}

	response, err := h.service.GetStudents(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)

	log.Info().
		Str("component", "students_handler").
		Msg("Students retrieved successfully")
}
