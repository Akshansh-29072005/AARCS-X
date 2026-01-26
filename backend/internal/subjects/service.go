package subjects

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateSubject(ctx context.Context, req CreateSubjectRequest) (*Subject, error) {
	entity := &SubjectEntity{
		Name:       req.Name,
		Code:       req.Code,
		SemesterId: req.SemesterId,
	}

	saved, err := s.repo.Create(ctx, entity)
	if err != nil {
		return nil, err
	}

	return &Subject{
		ID:         saved.ID,
		Name:       saved.Name,
		Code:       saved.Code,
		SemesterId: saved.SemesterId,
		CreatedAt:  saved.CreatedAt,
	}, nil
}

func (s *Service) GetSubjects(ctx context.Context, q GetSubjectRequest) (*GetSubjectResponse, error) {
	entities, err := s.repo.List(ctx, q)
	if err != nil {
		return nil, err
	}

	total, err := s.repo.Count(ctx, q)
	if err != nil {
		return nil, err
	}

	subjects := make([]SubjectListItem, 0, len(entities))
	for _, e := range entities {
		subjects = append(subjects, SubjectListItem{
			ID:         e.ID,
			Name:       e.Name,
			Code:       e.Code,
			SemesterId: e.SemesterId,
		})
	}

	return &GetSubjectResponse{
		Subjects: subjects,
		Total:    total,
	}, nil
}
