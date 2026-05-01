package auth

import (
	"context"

	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/errors"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/utils"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/users"
)

type Service struct {
	userService *users.Service
}

func NewService(userService *users.Service) *Service {
	return &Service{
		userService: userService,
	}
}

// RegisterUser service (user onboarding)
func (s *Service) RegisterUser(ctx context.Context, req CreateUserRequest) (string, error) {

	userId, err := s.userService.CreateUser(ctx, users.CreateUserRequest{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	})
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(userId)
	if err != nil {
		return "", errors.Internal("failed to generate token", err)
	}

	return token, nil
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (string, error) {

	user, err := s.userService.GetUserByEmail(ctx, req.Email, req.Password)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", errors.Internal("failed to generate token", err)
	}

	return token, nil
}