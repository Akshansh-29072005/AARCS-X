package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler wires HTTP requests to the auth service.
type Handler struct {
	service *Service
}

// NewHandler constructs an auth HTTP handler.
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Login authenticates a user and returns token pair.
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": ErrInvalidCredentials.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Refresh exchanges a refresh token for a new token pair.
func (h *Handler) Refresh(c *gin.Context) {
	var req RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.Refresh(c.Request.Context(), req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": ErrInvalidRefreshToken.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Logout revokes the provided refresh token.
func (h *Handler) Logout(c *gin.Context) {
	var req RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.RevokeRefreshToken(c.Request.Context(), req.RefreshToken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to revoke token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
