package students

import (
	"time"
)

type StudentEntity struct {
	ID            int
	Name          string
	Email         string
	Phone         string
	Password      string
	SemesterId    int
	DepartmentId  int
	InstitutionId int
	CreatedAt     time.Time
}
