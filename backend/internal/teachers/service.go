package teachers

import "context"

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
