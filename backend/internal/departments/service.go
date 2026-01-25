package departments

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateDepartment(ctx context.Context, req CreateDepartmentRequest) (*Department, error) {
	entity := &DepartmentEntity{
		Name:             req.Name,
		Code:             req.Code,
		HeadOfDepartment: req.HeadOfDepartment,
		InstitutionId:    req.InstitutionId,
	}

	saved, err := s.repo.Create(ctx, entity)
	if err != nil {
		return nil, err
	}

	return &Department{
		ID:               saved.ID,
		Name:             saved.Name,
		Code:             saved.Code,
		HeadOfDepartment: saved.HeadOfDepartment,
		InstitutionId:    saved.InstitutionId,
		CreatedAt:        saved.CreatedAt,
	}, nil
}

func (s *Service) GetDepartments(ctx context.Context, q GetDepartmentRequest) (*GetDepartmentResponse, error) {
	entities, err := s.repo.List(ctx, q)
	if err != nil {
		return nil, err
	}

	total, err := s.repo.Count(ctx, q)
	if err != nil {
		return nil, err
	}

	departments := make([]DepartmentListItem, 0, len(entities))
	for _, e := range entities {
		departments = append(departments, DepartmentListItem{
			ID:               e.ID,
			Name:             e.Name,
			Code:             e.Code,
			HeadOfDepartment: e.HeadOfDepartment,
			InstitutionId:    e.InstitutionId,
		})
	}

	return &GetDepartmentResponse{
		Departments: departments,
		Total:       total,
	}, nil
}
