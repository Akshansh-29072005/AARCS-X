package students

import "time"

type Student struct {
	ID            int       `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	Email         string    `json:"email" db:"email"`
	Phone         string    `json:"phone" db:"phone"`
	Password      string    `json:"password" db:"password"`
	SemesterId    int       `json:"semester" db:"semester_id"`
	DepartmentId  int       `json:"department" db:"department"`
	InstitutionId int       `json:"institution" db:"institution"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}
