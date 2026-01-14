package auth

import "time"

// User represents an authenticated user in the system.
type User struct {
	ID           int
	Email        string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
}

// RefreshToken models a persisted refresh token record.
type RefreshToken struct {
	ID        int
	UserID    int
	TokenHash string
	ExpiresAt time.Time
	CreatedAt time.Time
}
