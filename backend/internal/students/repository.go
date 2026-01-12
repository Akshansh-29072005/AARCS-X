package students

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, e *StudentEntity) (*StudentEntity, error) {
	query := `
		INSERT INTO students (first_name, last_name, email, phone, semester, branch, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,NOW(),NOW())
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRow(ctx, query,
		e.FirstName,
		e.LastName,
		e.Email,
		e.Phone,
		e.Semester,
		e.Branch,
	).Scan(&e.ID, &e.CreatedAt, &e.UpdatedAt)

	return e, err
}
