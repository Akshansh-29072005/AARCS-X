package institutes

import "time"

type InstitutionEntity struct {
	ID     	      int
	Name    	  string
	Code      	  string
	OfficialEmail string
	Address  	  string
	District	  string
	State    	  string
	Country 	  string
	CreatedAt	  time.Time
}

type AdminEntity struct {
	UserID         int
	InstitutionID  int  
	Role           string 
	CreatedAt      time.Time 
}

type GetByIDInstituteResponse struct {
	ID   int
	Name string
	Code string
}