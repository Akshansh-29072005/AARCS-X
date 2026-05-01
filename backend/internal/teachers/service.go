package teachers

import (
	"context"

	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/errors"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/utils"
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

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.Internal("failed to hash password", err)
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
		return nil, errors.FromPostgresError(err)
	}

	// Create user entry for authentication
	err = s.repo.CreateUser(ctx, req.Email, hashedPassword, saved.ID)
	if err != nil {
		return nil, errors.FromPostgresError(err)
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
		return nil, errors.FromPostgresError(err)
	}

	total, err := s.repo.Count(ctx, q)
	if err != nil {
		return nil, errors.FromPostgresError(err)
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

func (s *Service) GetTeacherByID(ctx context.Context, id int) (*GetByIDTeacherResponse, bool, error) {
	teacher, cacheHit, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, false, errors.FromPostgresError(err)
	}

	if teacher == nil {
		return nil, false, errors.NotFound("teacher not found", nil)
	}

	return teacher, cacheHit, nil
}