package teachers

import (
	"context"

	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/utlis"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateTeacher(ctx context.Context, req CreateTeacherRequest) (*Teacher, error) {

	hashedPassword, err := utlis.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	entity := &TeacherEntity{
		Name:         req.Name,
		Email:        req.Email,
		Phone:        req.Phone,
		Password:     hashedPassword,
		DepartmentId: req.DepartmentId,
		Designation:  req.Designation,
	}

	saved, err := s.repo.Create(ctx, entity)
	if err != nil {
		return nil, err
	}

	// Create user entry for authentication
	err = s.repo.CreateUser(ctx, req.Email, hashedPassword, saved.ID)
	if err != nil {
		return nil, err
	}

	return &Teacher{
		ID:           saved.ID,
		Name:         saved.Name,
		Email:        saved.Email,
		Phone:        saved.Phone,
		DepartmentId: saved.DepartmentId,
		Designation:  saved.Designation,
		CreatedAt:    saved.CreatedAt,
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
			ID:           e.ID,
			Name:         e.Name,
			DepartmentId: e.DepartmentId,
			Designation:  e.Designation,
		})
	}

	return &GetTeachersResponse{
		Teachers: teachers,
		Total:    total,
	}, nil
}
