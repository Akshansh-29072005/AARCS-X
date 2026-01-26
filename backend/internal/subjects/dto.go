package subjects

type CreateSubjectRequest struct {
	Name       string `json:"name" binding:"required"`
	Code       string `json:"code" binding:"required"`
	SemesterId int    `json:"semester_id" binding:"required"`
}

type SubjectListItem struct {
	ID         int    `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Code       string `json:"code" binding:"required"`
	SemesterId int    `json:"semester_id" binding:"required"`
}

type GetSubjectRequest struct {
	Name       string `form:"name"`
	Code       string `form:"code"`
	SemesterId int    `form:"semester_id" binding:"required"`
}

type GetSubjectResponse struct {
	Subjects []SubjectListItem `json:"subjects" binding:"required"`
	Total    int               `json:"total" binding:"required"`
}
