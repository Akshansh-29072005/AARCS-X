package students

type CreateStudentRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
	Semester  int    `json:"semester" binding:"required"`
	Branch    string `json:"branch" binding:"required"`
}

type StudentListItem struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Semester  int    `json:"semester"`
	Branch    string `json:"branch"`
}

type GetStudentsRequest struct {
	Branch   string `form:"branch"`
	Semester int    `form:"semester"`
}

type GetStudentsResponse struct {
	Students []StudentListItem `json:"students" binding:"required"`
	Total    int               `json:"total" binding:"required"`
}
