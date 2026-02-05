package auth

import (
	"context"

	"github.com/Akshansh-29072005/AARCS-X/backend/internal/utlis"
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
	token, err := GenerateToken(
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
	createInstitution func(ctx context.Context, name string) (int, error),
) (string, error) {

	// 1️⃣ Create institution
	institutionID, err := createInstitution(ctx, req.InstitutionName)
	if err != nil {
		return "", err
	}

	// 2️⃣ Hash password
	hashed, err := HashPassword(req.Password)
	if err != nil {
		return "", err
	}

	// 3️⃣ Create user
	userID, err := s.repo.CreateUser(
		ctx,
		req.Email,
		hashed,
		"institution_admin",
		&institutionID,
	)
	if err != nil {
		return "", err
	}

	// 4️⃣ Issue JWT
	return GenerateToken(
		userID,
		"institution_admin",
		&institutionID,
	)
}
