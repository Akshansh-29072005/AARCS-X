package teachers

import "time"

type TeacherEntity struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	Department  string
	Designation string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
