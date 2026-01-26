package subjects

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

func (r *Repository) Create(ctx context.Context, e *SubjectEntity) (*SubjectEntity, error) {
	query := `
		INSERT INTO subjects (name, code, semester_id, created_at)
		VALUES ($1,$2,$3,NOW())
		RETURNING id, created_at
	`

	err := r.db.QueryRow(ctx, query,
		e.Name,
		e.Code,
		e.SemesterId,
	).Scan(&e.ID, &e.CreatedAt)

	return e, err
}

func (r *Repository) List(ctx context.Context, q GetSubjectRequest) ([]Subject, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, name, code, semester_id
		FROM subjects
		WHERE ($1 = '' OR name = $1)
		AND ($2 = '' OR code = $2)
		AND ($3 = 0 OR semester_id = $3)`,
		q.Name, q.Code, q.SemesterId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var subjects []Subject
	for rows.Next() {
		var s Subject
		err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.Code,
			&s.SemesterId,
		)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return subjects, nil
}

func (r *Repository) Count(ctx context.Context, q GetSubjectRequest) (int, error) {
	var total int
	err := r.db.QueryRow(ctx,
		`SELECT COUNT(*) FROM subjects
			 WHERE ($1 = '' OR name = $1)
			 AND ($2 = '' OR code = $2)
			 AND ($3 = 0 OR semester_id = $3)`,
		q.Name, q.Code, q.SemesterId,
	).Scan(&total)
	return total, err
}
