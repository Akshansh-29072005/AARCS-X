package auth

type CreateUserRequest struct {
	Name     	string `json:"name" binding:"required"`
	Email    	string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required,e164"`
	Password 	string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string       `json:"token" binding:"required"`
}

type AuthenticatedUser struct {
	ID       int
	Role     string
	Email    string
	Password string
}