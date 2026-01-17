package teachers

import (
	"context"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTeacher(ctx context.Context, req CreateTeacherRequest) (*Teacher, error) {
	entity := &TeacherEntity{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Phone:       req.Phone,
		Department:  req.Department,
		Designation: req.Designation,
	}

	saved, err := s.repo.Create(ctx, entity)
	if err != nil {
		return nil, err
	}

	return &Teacher{
		ID:          saved.ID,
		FirstName:   saved.FirstName,
		LastName:    saved.LastName,
		Email:       saved.Email,
		Phone:       saved.Phone,
		Department:  saved.Department,
		Designation: saved.Designation,
		CreatedAt:   saved.CreatedAt,
		UpdatedAt:   saved.UpdatedAt,
	}, nil
}

func (s *Service) GetTeachers(ctx context.Context, q GetTeachersRequest) (*GetTeachersResponse, error) {
	entities, err := s.repo.List(ctx, q)
	if err != nil {
		return nil, err
	}

	total, err := s.repo.Count(ctx, q)
	if err != nil {
		return nil, err
	}

	teachers := make([]TeacherListItem, 0, len(entities))

	for _, e := range entities {
		teachers = append(teachers, TeacherListItem{
			ID:          e.ID,
			FirstName:   e.FirstName,
			LastName:    e.LastName,
			Department:  e.Department,
			Designation: e.Designation,
		})
	}

	return &GetTeachersResponse{
		Teachers: teachers,
		Total:    total,
	}, nil
}
