package teachers

type CreateTeacherRequest struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	Password     string `json:"password" binding:"required"`
	DepartmentId int    `json:"department_id" binding:"required"`
	Designation  string `json:"designation" binding:"required"`
}

type TeacherListItem struct {
	ID           int    `json:"id"`
	Name         string `json:"name" binding:"required"`
	DepartmentId int    `json:"department_id" binding:"required"`
	Designation  string `json:"designation" binding:"required"`
}

type GetTeachersRequest struct {
	DepartmentId int    `form:"department_id"`
	Designation  string `form:"designation"`
}

type GetTeachersResponse struct {
	Teachers []TeacherListItem `json:"teachers" binding:"required"`
	Total    int               `json:"total" binding:"required"`
}
