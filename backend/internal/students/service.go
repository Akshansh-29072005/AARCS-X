package students

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

func (s *Service) CreateStudent(ctx context.Context, req CreateStudentRequest) (*Student, error) {

	hashedPassword, err := utlis.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	entity := &StudentEntity{
		Name:          req.Name,
		Email:         req.Email,
		Phone:         req.Phone,
		Password:      hashedPassword,
		SemesterId:    req.SemesterId,
		DepartmentId:  req.DepartmentId,
		InstitutionId: req.InstitutionId,
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

	return &Student{
		ID:            saved.ID,
		Name:          saved.Name,
		Email:         saved.Email,
		Phone:         saved.Phone,
		SemesterId:    saved.SemesterId,
		DepartmentId:  saved.DepartmentId,
		InstitutionId: saved.InstitutionId,
		CreatedAt:     saved.CreatedAt,
	}, nil
}

func (s *Service) GetStudents(ctx context.Context, q GetStudentsRequest) (*GetStudentsResponse, error) {
	entities, err := s.repo.List(ctx, q)
	if err != nil {
		return nil, err
	}

	total, err := s.repo.Count(ctx, q)
	if err != nil {
		return nil, err
	}

	students := make([]StudentListItem, 0, len(entities))
	for _, e := range entities {
		students = append(students, StudentListItem{
			ID:            e.ID,
			Name:          e.Name,
			SemesterId:    e.SemesterId,
			DepartmentId:  e.DepartmentId,
			InstitutionId: e.InstitutionId,
		})
	}

	return &GetStudentsResponse{
		Students: students,
		Total:    total,
	}, nil
}
