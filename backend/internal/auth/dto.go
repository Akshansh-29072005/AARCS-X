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
