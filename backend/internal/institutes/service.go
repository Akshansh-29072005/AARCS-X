package institutes

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

func (s *Service) CreateInstitute(ctx context.Context, req CreateInstitutionRequest, userID int) (*Institute, error) {

	entity := &InstitutionEntity{
		Name:     	   req.Name,
		Code:          req.Code,
		OfficialEmail: req.Official_Email,
		Address:   	   req.Address,
		District:  	   req.District,
		State:     	   req.State,
		Country:  	   req.Country,
	}

	saved, err := s.repo.Create(ctx, entity, userID)
	if err != nil {
		return nil, errors.FromPostgresError(err)
	}

	return &Institute{
		ID:            saved.ID,
		Name:          saved.Name,
		Code:          saved.Code,
		OfficialEmail: saved.OfficialEmail,
		Address:       saved.Address,
		District:      saved.District,
		State:         saved.State,
		Country:       saved.Country,
		CreatedAt:     saved.CreatedAt,
	}, nil
}

func (s *Service) GetInstitutions(ctx context.Context, q GetInstitutionRequest) (*GetInstitutionsResponse, error) {
	entities, err := s.repo.List(ctx, q)
	if err != nil {
		return nil, errors.FromPostgresError(err)
	}

	total, err := s.repo.Count(ctx, q)
	if err != nil {
		return nil, errors.FromPostgresError(err)
	}

	institutions := make([]InstitutionListItem, 0, len(entities))
	for _, e := range entities {
		institutions = append(institutions, InstitutionListItem{
			ID:   e.ID,
			Name: e.Name,
			Code: e.Code,
		})
	}

	return &GetInstitutionsResponse{
		Institutions: institutions,
		Total:        total,
	}, nil
}

func (s *Service) GetInstitutionByID(ctx context.Context, id int) (*GetByIDInstituteResponse, bool, error) {
	institution, cacheHit, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, false, errors.FromPostgresError(err)
	}

	if institution == nil {
		return nil, false, errors.NotFound("institution not found", nil)
	}
	
	return institution, cacheHit, nil
}