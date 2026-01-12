package students

import (
	"time"
)

type StudentEntity struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Semester  int
	Branch    string
	CreatedAt time.Time
	UpdatedAt time.Time
}