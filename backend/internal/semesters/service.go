package semesters

import (
	"context"

	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/errors"
)

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
		return nil, errors.FromPostgresError(err)
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
		return nil, errors.FromPostgresError(err)
	}

	total, err := s.repo.Count(ctx, q)
	if err != nil {
		return nil, errors.FromPostgresError(err)
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

func (s *Service) GetSemesterByID(ctx context.Context, id int) (*GetByIDSemesterResponse, bool, error) {
	semester, cacheHit, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, false, errors.FromPostgresError(err)
	}

	if semester == nil {
		return nil, false, errors.NotFound("semester not found", nil)
	}

	return semester, cacheHit, nil
}