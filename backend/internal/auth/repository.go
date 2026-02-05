package auth

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

type User struct {
	ID          int
	Email       string
	Password    string
	Role        string
	ReferenceID *int
}

// GetByEmail fetches user by email (for login)
func (r *Repository) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, email, password, role, reference_id
		FROM users
		WHERE email = $1
	`

	var u User
	err := r.db.QueryRow(ctx, query, email).Scan(
		&u.ID,
		&u.Email,
		&u.Password,
		&u.Role,
		&u.ReferenceID,
	)

	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &u, nil
}

// CreateUser creates institution admin user
func (r *Repository) CreateUser(
	ctx context.Context,
	email, hashedPassword, role string,
	refID *int,
) (int, error) {

	query := `
		INSERT INTO users (email, password, role, reference_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var userID int
	err := r.db.QueryRow(
		ctx,
		query,
		email,
		hashedPassword,
		role,
		refID,
	).Scan(&userID)

	return userID, err
}
