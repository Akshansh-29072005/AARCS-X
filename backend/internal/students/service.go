package students

import (
	"context"

	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/database"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/errors"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/utlis"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	repo *Repository
	pool *pgxpool.Pool
}

func NewService(repo *Repository, pool *pgxpool.Pool) *Service {
	return &Service{
		repo: repo,
		pool: pool,
	}
}

func (s *Service) CreateStudent(ctx context.Context, req CreateStudentRequest) (*Student, error) {

	var student *Student

	err := database.WithTransaction(ctx, s.pool, func(tx database.DBTX) error{
		repo := NewRepository(tx)

		hashedPassword, err := utlis.HashPassword(req.Password)
		if  err != nil {
			return errors.Internal("failed to hash password", err)
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

		saved, err := repo.Create(ctx, entity)
		if err != nil {
			return errors.FromPostgresError(err)
		}

		// Create user entry for authentication
		err = repo.CreateUser(ctx, req.Email, hashedPassword, saved.ID)
		if err != nil {
			return errors.FromPostgresError(err)
		}

		student = &Student{
			ID:            saved.ID,
			Name:          saved.Name,
			Email:         saved.Email,
			Phone:         saved.Phone,
			SemesterId:    saved.SemesterId,
			DepartmentId:  saved.DepartmentId,
			InstitutionId: saved.InstitutionId,
			CreatedAt:     saved.CreatedAt,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return student, nil
}

func (s *Service) GetStudents(ctx context.Context, q GetStudentsRequest) (*GetStudentsResponse, error) {
	entities, err := s.repo.List(ctx, q)
	if err != nil {
		return nil, errors.FromPostgresError(err)
	}

	total, err := s.repo.Count(ctx, q)
	if err != nil {
		return nil, errors.FromPostgresError(err)
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
