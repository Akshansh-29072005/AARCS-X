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

type GetByIDStudentResponse struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	SemesterId    int       `json:"semester_id"`
	DepartmentId  int       `json:"department_id"`
	InstitutionId int       `json:"institution_id"`
}