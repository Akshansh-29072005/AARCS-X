package teachers

type CreateTeacherRequest struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	Department  string `json:"department" binding:"required"`
	Designation string `json:"designation" binding:"required"`
}

type TeacherListItem struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Department  string `json:"department" binding:"required"`
	Designation string `json:"designation" binding:"required"`
}

type GetTeachersRequest struct {
	Department  string `json:"department" binding:"required"`
	Designation string `json:"designation" binding:"required"`
}

type GetStudentsRequest struct {
	Teachers []TeacherListItem `json:"teachers" binding:"required"`
	Total    int               `json:"total" binding:"required"`
}
