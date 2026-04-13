package users

type CreateUserRequest struct {
	Name     	string `json:"name" binding:"required"`
	Email    	string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required,e164"`
	Password 	string `json:"password" binding:"required,min=8"`
}