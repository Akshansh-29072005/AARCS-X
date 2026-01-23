package students

import "time"

type Student struct {
	ID            int       `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	Email         string    `json:"email" db:"email"`
	Phone         string    `json:"phone" db:"phone"`
	Password      string    `json:"password" db:"password"`
	SemesterId    int       `json:"semester_id" db:"semester_id"`
	DepartmentId  int       `json:"department_id" db:"department_id"`
	InstitutionId int       `json:"institution_id" db:"institution_id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}
