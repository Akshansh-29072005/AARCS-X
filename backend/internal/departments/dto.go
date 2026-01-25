package departments

type CreateDepartmentRequest struct {
	Name             string `json:"name" binding:"required"`
	Code             string `json:"code" binding:"required"`
	HeadOfDepartment string `json:"head_of_department" binding:"required"`
	InstitutionId    int    `json:"institution_id" binding:"required"`
}

type DepartmentListItem struct {
	ID               int    `json:"id" binding:"required"`
	Name             string `json:"name" binding:"required"`
	Code             string `json:"code" binding:"required"`
	HeadOfDepartment string `json:"head_of_department" binding:"required"`
	InstitutionId    int    `json:"institution_id" binding:"required"`
}

type GetDepartmentRequest struct {
	Name             string `form:"name"`
	Code             string `form:"code"`
	HeadOfDepartment string `form:"head_of_department"`
	InstitutionId    int    `form:"institution_id"`
}

type GetDepartmentResponse struct {
	Departments []DepartmentListItem `json:"departments" binding:"required"`
	Total       int                  `json:"total" binding:"required"`
}
