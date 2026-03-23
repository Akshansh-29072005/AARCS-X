package institutes

import "time"

type InstitutionEntity struct {
	ID        int
	Name      string
	Code      string
	Password  string
	CreatedAt time.Time
}

type GetByIDInstituteResponse struct {
	ID   int
	Name string
	Code string
}