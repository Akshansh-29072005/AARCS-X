package semesters

import "time"

type Semester struct {
	ID           int       `json:"id" db:"id"`
	Number       int       `json:"number" db:"number"`
	DepartmentId int       `json:"department_id" db:"department_id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
