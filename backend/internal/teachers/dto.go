package teachers

type CreateTeacherRequest struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNo     string `json:"phone_no" binding:"required"`
	Department  string `json:"department" binding:"required"`
	Designation string `json:"designation" binding:"required"`
}