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

func (r *Repository) List(ctx context.Context, q GetStudentsRequest) ([]Student, error) {
	rows, err := r.db.Query(ctx,
	   `SELECT id, first_name, last_name, semester, branch
		FROM students
		WHERE ($1 = '' OR branch = $1)
		AND ($2 = 0 OR semester = $2)`,
		q.Branch, q.Semester,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []Student
	for rows.Next() {
		var s Student
		err := rows.Scan(
			&s.ID,
			&s.FirstName,
			&s.LastName,
			&s.Semester,
			&s.Branch,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (r *Repository) Count(ctx context.Context, q GetStudentsRequest) (int, error) {
	var total int
	err := r.db.QueryRow(ctx,
			`SELECT COUNT(*) FROM students
			 WHERE ($1 = '' OR branch = $1)
			 AND ($2 = 0 OR semester = $2)`,
			 q.Branch, q.Semester,
			 ).Scan(&total)
	return total, err
}