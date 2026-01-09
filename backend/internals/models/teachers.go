package models

import "time"

type Teacher struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Phone       string `json:"phone" db:"phone"`
	Department  string `json:"department" db:"department"`
	Designation string `json:"designation" db:"designation"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}