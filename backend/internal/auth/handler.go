package auth

import "time"

type UserEntity struct {
	ID           int
	Email        string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
}

type RefreshTokenEntity struct {
	ID        int
	UserID    int
	TokenHash string
	ExpiresAt time.Time
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(200, resp)
}
