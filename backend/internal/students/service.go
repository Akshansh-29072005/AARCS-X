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
			ID:        e.ID,
			FirstName: e.FirstName,
			LastName:  e.LastName,
			Semester:  e.Semester,
			Branch:    e.Branch,
		})
	}

	return &GetStudentsResponse{
		Students: students,
		Total: total,
	}, nil
}
