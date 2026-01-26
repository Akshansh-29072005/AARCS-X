package semesters

import "time"

type SemesterEntity struct {
	ID           int
	Number       int
	DepartmentId int
	CreatedAt    time.Time
}
