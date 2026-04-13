package users

import (
	"context"
	"strings"

	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/errors"
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

func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) (int, error) {

	// CleanedEmail and CleanedPhone are used to ensure that the email and phone number are stored in a consistent format.
	cleanedEmail := strings.ToLower(strings.TrimSpace(req.Email))
	cleanedPhone := strings.TrimSpace(req.PhoneNumber)


	hashedPassword, err := utlis.HashPassword(req.Password)
	if err != nil {
		return 0, errors.Internal("failed to hash password", err)
	}

	entity := &UserEntity{
		Name:        req.Name,
		Email:       cleanedEmail,
		PhoneNumber: cleanedPhone,
		Password:    hashedPassword,
	}
	
	saved, err := s.repo.CreateUser(ctx, entity)
	if err != nil {
		return 0, errors.FromPostgresError(err)
	}

	return saved.ID, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, email string, password string) (*UserEntity, error) {

	// CleanedEmail is used to ensure that the email is stored in a consistent format.
	cleanedEmail := strings.ToLower(strings.TrimSpace(email))

	user, err := s.repo.GetUserByEmail(ctx, cleanedEmail)

	if err != nil {
		return nil, errors.FromPostgresError(err)
	} else if user == nil {
		return nil, errors.NotFound("user not found", nil)
	}

	isMatch, err := utlis.ComparePasswords(user.Password, password)
	if err != nil {
		return nil, errors.Internal("failed to compare passwords", err)
	}
	if !isMatch {
		return nil, errors.Unauthorized("invalid email or password", nil)
	}

	return user, nil
}