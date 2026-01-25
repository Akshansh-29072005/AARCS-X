package institutes

type CreateInstitutionRequest struct {
	Name     string `json:"name" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required"`
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
