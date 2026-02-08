package auth

type UserResponse struct {
	Id    int    `json:"id" binding:"required"`
	Role  string `json:"role" binding:"required"`
	Email string `json:"email" binding:"required"`
}

// Login Info

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string       `json:"token" binding:"required"`
	User  UserResponse `json:"user" binding:"required"`
}

type AuthenticatedUser struct {
	ID       int
	Role     string
	Email    string
	Password string
}

type RegisterRequest struct {
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required"`
	InstitutionName string `json:"institution" binding:"required"`
	ReferenceId     int    `json:"reference_id,omitempty"`
}
