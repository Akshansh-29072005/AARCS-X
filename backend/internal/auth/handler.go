package auth

import (
	"context"
	"net/http"

	"github.com/Akshansh-29072005/AARCS-X/backend/internal/institutes"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service            *Service
	institutionService institutes.Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Login Handler
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// RegisterInstitution Handler
func (h *Handler) RegisterInstitution(
	createInstitution func(ctx context.Context, name string) (int, error),
) gin.HandlerFunc {

	return func(c *gin.Context) {
		var req RegisterRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := h.service.RegisterInstitution(
			c.Request.Context(),
			req,
			h.institutionService.Create,
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"token": token})
	}
}

// Me route handler
func (h *Handler) Me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user_id": c.GetInt("user_id"),
		"role":    c.GetString("role"),
		"ref_id":  c.MustGet("ref_id"),
	})
}
