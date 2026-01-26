package subjects

import "time"

type SubjectEntity struct {
	ID         int
	Name       string
	Code       string
	SemesterId int
	CreatedAt  time.Time
}
