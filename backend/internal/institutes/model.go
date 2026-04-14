package institutes

import "time"

type Institute struct {
	ID       	  int       `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	Code       	  string    `json:"code" db:"code"`
	OfficialEmail string `json:"official_email" db:"official_email"`
	Address       string    `json:"address" db:"address"`
	District      string    `json:"district" db:"district"`
	State         string    `json:"state" db:"state"`
	Country       string    `json:"country" db:"country"`
	CreatedAt	  time.Time `json:"created_at" db:"created_at"`
}
