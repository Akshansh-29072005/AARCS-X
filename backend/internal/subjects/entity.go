package subjects

import "time"

type SubjectEntity struct {
	ID         int
	Name       string
	Code       string
	SemesterId int
	CreatedAt  time.Time
}

type GetByIDSubjectResponse struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Code       string    `json:"code"`
	SemesterId int       `json:"semester_id"`
}