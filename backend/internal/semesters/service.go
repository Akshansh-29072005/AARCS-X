package semesters

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateSemesters(ctx context.Context, req CreateSemesterRequest) (*Semester, error) {
	entity := &SemesterEntity{
		Number:       req.Number,
		DepartmentId: req.DepartmentId,
	}

	saved, err := s.repo.Create(ctx, entity)
	if err != nil {
		return nil, err
	}

	return &Semester{
		ID:           saved.ID,
		Number:       saved.Number,
		DepartmentId: saved.DepartmentId,
		CreatedAt:    saved.CreatedAt,
	}, nil
}

func (s *Service) GetSemesters(ctx context.Context, q GetSemestersRequest) (*GetSemestersResponse, error) {
	entities, err := s.repo.List(ctx, q)
	if err != nil {
		return nil, err
	}

	total, err := s.repo.Count(ctx, q)
	if err != nil {
		return nil, err
	}

	semesters := make([]SemesterListItem, 0, len(entities))
	for _, e := range entities {
		semesters = append(semesters, SemesterListItem{
			ID:           e.ID,
			Number:       e.Number,
			DepartmentId: e.DepartmentId,
		})
	}

	return &GetSemestersResponse{
		Semesters: semesters,
		Total:     total,
	}, nil
}
