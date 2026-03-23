package departments

import "time"

type DepartmentEntity struct {
	ID               int
	Name             string
	Code             string
	HeadOfDepartment string
	InstitutionId    int
	CreatedAt        time.Time
}

type GetByIDDepartmentResponse struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Code             string `json:"code"`
	HeadOfDepartment string `json:"head_of_department"`
	InstitutionId    int    `json:"institution_id"`
}