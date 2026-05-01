package users

import (
	"context"
	"strconv"
	"strings"
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
	WITH inserted_user AS (
		INSERT INTO users (name, email, phone, password_hash, created_at)
		VALUES ($1, $2, $3, $4, NOW())
		RETURNING id, created_at
	),
	inserted_role AS (
		INSERT INTO roles (user_id, role, created_at)
		SELECT id, 'user', NOW() FROM inserted_user
		RETURNING id, created_at
	)
	SELECT 
		id,
		created_at
	FROM inserted_user;
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

func (r *Repository) GetRolesByID(ctx context.Context, id int) ([]string, error) {

	// Check cache first
	cacheKey := "user_roles:" + strconv.Itoa(id)
	cachedRoles, err := r.rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		return strings.Split(cachedRoles, ","), nil
	}

	// If cache miss, query database
	query := `
		SELECT role
		FROM roles
		WHERE user_id = $1
	`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []string
	for rows.Next() {
		var role string
		if err := rows.Scan(&role); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Cache the roles in Redis
	cacheValue := strings.Join(roles, ",")
	_ = r.rdb.Set(context.Background(), cacheKey, cacheValue, UsersCacheTTL).Err()

	return roles, nil
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