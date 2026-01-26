package semesters

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

func (r *Repository) Create(ctx context.Context, e *SemesterEntity) (*SemesterEntity, error) {
	query := `
		INSERT INTO semesters (number, department_id, created_at)
		VALUES ($1,$2,NOW())
		RETURNING id, created_at
	`

	err := r.db.QueryRow(ctx, query,
		e.Number,
		e.DepartmentId,
	).Scan(&e.ID, &e.CreatedAt)

	return e, err
}

func (r *Repository) List(ctx context.Context, q GetSemestersRequest) ([]Semester, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, number, department_id
		FROM semesters
		WHERE ($1 = 0 OR number = $1)
		AND ($2 = 0 OR department_id = $2)`,
		q.Number, q.DepartmentId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var semesters []Semester
	for rows.Next() {
		var s Semester
		err := rows.Scan(
			&s.ID,
			&s.Number,
			&s.DepartmentId,
		)
		if err != nil {
			return nil, err
		}
		semesters = append(semesters, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return semesters, nil
}

func (r *Repository) Count(ctx context.Context, q GetSemestersRequest) (int, error) {
	var total int
	err := r.db.QueryRow(ctx,
		`SELECT COUNT(*) FROM semesters
			 WHERE ($1 = 0 OR number = $1)
			 AND ($2 = 0 OR department_id = $2)
			 GROUP BY number, department_id`,
		q.Number, q.DepartmentId,
	).Scan(&total)
	return total, err
}
