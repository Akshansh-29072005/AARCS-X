package institutes

type CreateInstitutionRequest struct {
	Name  		   string `json:"name" binding:"required"`
	Code     	   string `json:"code" binding:"required"`
	Official_Email string `json:"official_email" binding:"required"`
	Address        string `json:"address" binding:"required"`
	District       string `json:"district" binding:"required"`
	State          string `json:"state" binding:"required"`
	Country        string `json:"country" binding:"required"`
}

type InstitutionListItem struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type GetInstitutionRequest struct {
	Name string `form:"name"`
	Code string `form:"code"`
}

type GetInstitutionsResponse struct {
	Institutions []InstitutionListItem `json:"institutions" binding:"required"`
	Total        int                   `json:"total" binding:"required"`
}
