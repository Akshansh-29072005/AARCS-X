package subjects

import "time"

type Subject struct {
	ID         int       `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Code       string    `json:"code" db:"code"`
	SemesterId int       `json:"semester_id" db:"semester_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
