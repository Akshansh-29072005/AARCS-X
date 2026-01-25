package institutes

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateInstitution(ctx context.Context, req CreateInstitutionRequest) (*Institute, error) {
	entity := &InstitutionEntity{
		Name:     req.Name,
		Code:     req.Code,
		Password: req.Password,
	}

	saved, err := s.repo.Create(ctx, entity)
	if err != nil {
		return nil, err
	}

	return &Institute{
		ID:        saved.ID,
		Name:      saved.Name,
		Code:      saved.Code,
		CreatedAt: saved.CreatedAt,
	}, nil
}

func (s *Service) GetInstitutions(ctx context.Context, q GetInstitutionRequest) (*GetInstitutionsResponse, error) {
	entities, err := s.repo.List(ctx, q)
	if err != nil {
		return nil, err
	}

	total, err := s.repo.Count(ctx, q)
	if err != nil {
		return nil, err
	}

	institutions := make([]InstitutionListItem, 0, len(entities))
	for _, e := range entities {
		institutions = append(institutions, InstitutionListItem{
			ID:   e.ID,
			Name: e.Name,
			Code: e.Code,
		})
	}

	return &GetInstitutionsResponse{
		Institutions: institutions,
		Total:        total,
	}, nil
}
