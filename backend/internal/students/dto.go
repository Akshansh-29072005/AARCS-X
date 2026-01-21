package students

type CreateStudentRequest struct {
	Name          string `json:"name" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
	Phone         string `json:"phone" binding:"required"`
	Password      string `json:"password" binding:"required"`
	SemesterId    int    `json:"semester_id" binding:"required"`
	DepartmentId  int    `json:"department_id" binding:"required"`
	InstitutionId int    `json:"institution_id" binding:"required"`
}

type StudentListItem struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	SemesterId    int    `json:"semester_id"`
	DepartmentId  int    `json:"department_id"`
	InstitutionId int    `json:"institution_id"`
}

type GetStudentsRequest struct {
	SemesterId    int `form:"semester_id"`
	DepartmentId  int `form:"department_id"`
	InstitutionId int `form:"institution_id"`
}

type GetStudentsResponse struct {
	Students []StudentListItem `json:"students" binding:"required"`
	Total    int               `json:"total" binding:"required"`
}
