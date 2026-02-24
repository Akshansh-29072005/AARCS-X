package auth

import (
	"context"

	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/utlis"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// Login service
func (s *Service) Login(ctx context.Context, req LoginRequest) (string, error) {

	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return "", err
	}

	// Compare password (bcrypt)
	err = utlis.ComparePasswords(user.Password, req.Password)
	if err != nil {
		return "", err
	}

	// Generate JWT
	token, err := utlis.GenerateToken(
		user.ID,
		user.Role,
		user.ReferenceID,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}

// RegisterInstitution service (institution admin onboarding)
func (s *Service) RegisterInstitution(
	ctx context.Context,
	req RegisterRequest,
	createInstitution func(ctx context.Context, name, code, password string) (int, error),
) (string, error) {

	// Create institution
	institutionID, err := createInstitution(ctx, req.InstitutionName, req.InstitutionCode, req.Password)
	if err != nil {
		return "", err
	}

	// Hash password
	hashed, err := utlis.HashPassword(req.Password)
	if err != nil {
		return "", err
	}

	// Create user
	userID, err := s.repo.CreateUser(
		ctx,
		req.Email,
		hashed,
		"institution",
		&institutionID,
	)
	if err != nil {
		return "", err
	}

	// Issue JWT
	return utlis.GenerateToken(
		userID,
		"institution",
		&institutionID,
	)
}
