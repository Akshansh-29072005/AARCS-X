package semesters

import "time"

type SemesterEntity struct {
	ID           int
	Number       int
	DepartmentId int
	CreatedAt    time.Time
}

type GetByIDSemesterResponse struct {
	ID           int       `json:"id"`
	Number       int       `json:"number"`
	DepartmentId int       `json:"department_id"`
}
