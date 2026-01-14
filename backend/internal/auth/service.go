package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo       Repository
	secret     string
	accessTTL  time.Duration
	refreshTTL time.Duration
	bcryptCost int
}

var (
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)

// NewService constructs an auth service with configured token TTLs.
func NewService(
	repo Repository,
	secret string,
	accessTTL, refreshTTL time.Duration,
) *Service {
	return &Service{
		repo:       repo,
		secret:     secret,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
		bcryptCost: bcrypt.DefaultCost,
	}
}

// Login validates credentials and issues a new token pair.
func (s *Service) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := ComparePassword(user.PasswordHash, req.Password); err != nil {
		return nil, ErrInvalidCredentials
	}

	accessToken, accessExp, err :=
		GenerateAccessToken(user.ID, user.Role, s.secret, s.accessTTL)
	if err != nil {
		return nil, err
	}

	refreshToken, _, err := s.issueRefreshToken(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(time.Until(accessExp).Seconds()),
	}, nil
}

// Refresh exchanges a valid refresh token for a new token pair.
func (s *Service) Refresh(ctx context.Context, token string) (*RefreshResponse, error) {
	hashed := hashToken(token)

	stored, err := s.repo.GetRefreshToken(ctx, hashed)
	if err != nil || time.Now().After(stored.ExpiresAt) {
		_ = s.repo.DeleteRefreshToken(ctx, hashed)
		return nil, ErrInvalidRefreshToken
	}

	// 🔐 Rotate first to prevent reuse
	if err := s.repo.DeleteRefreshToken(ctx, hashed); err != nil {
		return nil, ErrInvalidRefreshToken
	}

	user, err := s.repo.GetUserByID(ctx, stored.UserID)
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	accessToken, accessExp, err :=
		GenerateAccessToken(user.ID, user.Role, s.secret, s.accessTTL)
	if err != nil {
		return nil, err
	}

	newRefreshToken, _, err := s.issueRefreshToken(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return &RefreshResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    int64(time.Until(accessExp).Seconds()),
	}, nil
}

// RevokeRefreshToken removes a refresh token, used for logout flows.
func (s *Service) RevokeRefreshToken(ctx context.Context, token string) error {
	return s.repo.DeleteRefreshToken(ctx, hashToken(token))
}

// HashPassword produces a bcrypt hash for a password.
func (s *Service) HashPassword(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), s.bcryptCost)
	return string(b), err
}

// ComparePassword checks a plaintext password against a bcrypt hash.
func ComparePassword(hash, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}

func (s *Service) issueRefreshToken(ctx context.Context, userID int) (string, time.Time, error) {
	// 32 bytes → 64 hex chars (strong enough)
	refreshToken, err := generateRandomTokenBytes(32)
	if err != nil {
		return "", time.Time{}, err
	}

	expiresAt := time.Now().Add(s.refreshTTL)

	record := RefreshToken{
		UserID:    userID,
		TokenHash: hashToken(refreshToken),
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
	}

	if err := s.repo.SaveRefreshToken(ctx, record); err != nil {
		return "", time.Time{}, err
	}

	return refreshToken, expiresAt, nil
}

// generateRandomTokenBytes generates cryptographically secure tokens.
func generateRandomTokenBytes(byteLen int) (string, error) {
	b := make([]byte, byteLen)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func hashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}
