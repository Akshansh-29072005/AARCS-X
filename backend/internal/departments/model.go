package departments

import "time"

type Department struct {
	ID               int       `json:"id" db:"id"`
	Name             string    `json:"name" db:"name"`
	Code             string    `json:"code" db:"code"`
	HeadOfDepartment string    `json:"head_of_department" db:"head_of_department"`
	InstitutionId    int       `json:"institution_id" db:"institution_id"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
}
