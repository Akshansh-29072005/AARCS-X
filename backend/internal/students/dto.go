package students

type CreateStudentRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
	Semester  int    `json:"semester" binding:"required"`
	Branch    string `json:"branch" binding:"required"`
}