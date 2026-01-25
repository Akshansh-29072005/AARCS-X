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
