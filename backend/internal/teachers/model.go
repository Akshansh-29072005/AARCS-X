package teachers

import "time"

type Teacher struct {
	ID           int       `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Email        string    `json:"email" db:"email"`
	Phone        string    `json:"phone" db:"phone"`
	DepartmentId int       `json:"department_id" db:"department_id"`
	Designation  string    `json:"designation" db:"designation"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
