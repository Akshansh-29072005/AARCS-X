package auth

import (
	"context"
	"net/http"

	"github.com/Akshansh-29072005/AARCS-X/backend/internal/institutes"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/middleware"
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

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "auth_handler").
		Msg("Received login request")

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

	log.Info().
		Str("component", "auth_handler").
		Msg("Login successful")
}

// RegisterInstitution Handler
func (h *Handler) RegisterInstitution(c *gin.Context) {

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "auth_handler").
		Msg("Received institution registration request")
	
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 🔹 Adapter function
	createInstitution := func(ctx context.Context, name, code, password string) (int, error) {
		inst, err := h.institutionService.CreateInstitution(
			ctx,
			institutes.CreateInstitutionRequest{
				Name:     name,
				Code:     code,
				Password: password,
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

	log.Info().
		Str("component", "auth_handler").
		Msg("Institution registered successfully")
}

// Me route handler
func (h *Handler) Me(c *gin.Context) {

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "auth_handler").
		Msg("Received request for user info")
	
	c.JSON(http.StatusOK, gin.H{
		"user_id": c.GetInt("user_id"),
		"role":    c.GetString("role"),
		"ref_id":  c.MustGet("ref_id"),
	})

	log.Info().
		Str("component", "auth_handler").
		Msg("User info retrieved successfully")
}
