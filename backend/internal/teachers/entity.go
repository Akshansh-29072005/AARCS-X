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
