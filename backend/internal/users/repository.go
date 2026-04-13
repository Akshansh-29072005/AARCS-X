package users

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

const UsersCacheTTL = time.Minute * 10

type Repository struct {
	db *pgxpool.Pool
	rdb *redis.Client
}

func NewRepository(db *pgxpool.Pool, rdb *redis.Client) *Repository {
	return &Repository{
		db: db,
		rdb: rdb,
	}
}

func (r *Repository) CreateUser(ctx context.Context, u *UserEntity) (*UserEntity, error) {
	query := `
		INSERT INTO users (name, email, phone, password_hash, created_at)
		VALUES ($1, $2, $3, $4, NOW())
		RETURNING id, created_at
	`

	err := r.db.QueryRow(ctx, query,
		u.Name,
		u.Email,
		u.PhoneNumber,
		u.Password,
	).Scan(&u.ID, &u.CreatedAt)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*UserEntity, error) {
	query := `
		SELECT id, name, email, phone, password_hash, created_at
		FROM users
		WHERE email = $1
	`

	u := &UserEntity{}
	err := r.db.QueryRow(ctx, query, email).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.PhoneNumber,
		&u.Password,
		&u.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return u, nil
}