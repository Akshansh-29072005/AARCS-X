package institutes

type CreateInstitution struct {
	Name string `form:"name"`
	Code string `form:"code"`
}

type GetInstitutions struct {
	Name string `form:"name"`
	Code string `form:"code"`
}
