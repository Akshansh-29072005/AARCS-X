package students

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateStudent(ctx context.Context, req CreateStudentRequest) (*Student, error) {
	entity := &StudentEntity{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Semester:  req.Semester,
		Branch:    req.Branch,
	}

	saved, err := s.repo.Create(ctx, entity)
	if err != nil {
		return nil, err
	}

	return &Student{
		ID:        saved.ID,
		FirstName: saved.FirstName,
		LastName:  saved.LastName,
		Email:     saved.Email,
		Phone:     saved.Phone,
		Semester:  saved.Semester,
		Branch:    saved.Branch,
		CreatedAt: saved.CreatedAt,
		UpdatedAt: saved.UpdatedAt,
	}, nil
}
