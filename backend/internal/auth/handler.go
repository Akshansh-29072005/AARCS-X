package auth

import (
	"context"
	"net/http"

	"github.com/Akshansh-29072005/AARCS-X/backend/internal/institutes"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service            *Service
	institutionService *institutes.Service
}

func NewHandler(service *Service, institutionService *institutes.Service) *Handler {
	return &Handler{
		service:            service,
		institutionService: institutionService,
	}
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
func (h *Handler) RegisterInstitution(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 🔹 Adapter function
	createInstitution := func(ctx context.Context, name string) (int, error) {
		inst, err := h.institutionService.CreateInstitution(
			ctx,
			institutes.CreateInstitutionRequest{
				Name: name,
			},
		)
		if err != nil {
			return 0, err
		}
		return inst.ID, nil
	}

	token, err := h.service.RegisterInstitution(
		c.Request.Context(),
		req,
		createInstitution,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}

// Me route handler
func (h *Handler) Me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user_id": c.GetInt("user_id"),
		"role":    c.GetString("role"),
		"ref_id":  c.MustGet("ref_id"),
	})
}
