package semesters

type CreateSemesterRequest struct {
	Number       int `json:"number" binding:"required"`
	DepartmentId int `json:"department_id" binding:"required"`
}

type SemesterListItem struct {
	ID           int `json:"id" binding:"required"`
	Number       int `json:"number" binding:"required"`
	DepartmentId int `json:"department_id" binding:"required"`
}

type GetSemestersRequest struct {
	Number       int `form:"number"`
	DepartmentId int `form:"department_id"`
}

type GetSemestersResponse struct {
	Semesters []SemesterListItem `json:"semesters" binding:"required"`
	Total     int                `json:"total" binding:"required"`
}
