package auth

import (
	"net/http"

	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/errors"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service         *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// RegisterUser Handler
func (h *Handler) RegisterUser(c *gin.Context) {

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "auth_handler").
		Msg("Received request to register user")

	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.BadRequest("invalid request body", err))
		return
	}

	token, err := h.service.RegisterUser(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})

	log.Info().
		Str("component", "auth_handler").
		Msg("User registered successfully")
}

func (h *Handler) Login(c *gin.Context) {

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "auth_handler").
		Msg("Received request to login user")

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.BadRequest("invalid request body", err))
		return
	}

	token, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

	log.Info().
		Str("component", "auth_handler").
		Msg("User logged in successfully")
}

// Me Handler - returns user info based on token
func (h *Handler) Me(c *gin.Context) {

	log := middleware.GetLogger(c)
	log.Info().
		Str("component", "auth_handler").
		Msg("Received request for user info")
	
	c.JSON(http.StatusOK, gin.H{
		"user_id": c.GetInt("user_id"),
	})

	log.Info().
		Str("component", "auth_handler").
		Msg("User info retrieved successfully")
}
