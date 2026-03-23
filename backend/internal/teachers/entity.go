package teachers

import (
	"time"
)

type TeacherEntity struct {
	ID           int
	Name         string
	Email        string
	Phone        string
	Password     string
	DepartmentId int
	Designation  string
	CreatedAt    time.Time
}

type GetByIDTeacherResponse struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	DepartmentId int       `json:"department_id"`
	Designation  string    `json:"designation"`
}