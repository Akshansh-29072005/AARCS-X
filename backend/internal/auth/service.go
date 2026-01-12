package auth

import "golang.org/x/crypto/bcrypt"

type Service struct {
	repo   *Repository
	secret string
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := ComparePassword(user.PasswordHash, req.Password); err != nil {
		return nil, ErrInvalidCredentials
	}

	access, _ := GenerateJWT(user.ID, user.Role, s.secret)
	refresh := GenerateRandomToken() // crypto/rand

	s.repo.SaveRefreshToken(ctx, user.ID, refresh)

	return &LoginResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}


func HashPassword(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), 12)
	return string(b), err
}

func ComparePassword(hash, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}
