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
		INSERT INTO students (name, email, phone, password, semester_id, department_id, institution_id, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,NOW())
		RETURNING id, created_at
	`

	err := r.db.QueryRow(ctx, query,
		e.Name,
		e.Email,
		e.Phone,
		e.Password,
		e.SemesterId,
		e.DepartmentId,
		e.InstitutionId,
	).Scan(&e.ID, &e.CreatedAt)

	return e, err
}

func (r *Repository) List(ctx context.Context, q GetStudentsRequest) ([]Student, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, name, semester_id, department_id, institution_id
		FROM students
		WHERE ($1 = 0 OR semester_id = $1)
		AND ($2 = 0 OR department_id = $2)
		AND ($3 = 0 OR institution_id = $3)`,
		q.SemesterId, q.DepartmentId, q.InstitutionId,
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
			&s.Name,
			&s.SemesterId,
			&s.DepartmentId,
			&s.InstitutionId,
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
			 WHERE ($1 = 0 OR semester_id = $1)
			 AND ($2 = 0 OR department_id = $2)
			 AND ($3 = 0 OR institution_id = $3)`,
		q.SemesterId, q.DepartmentId, q.InstitutionId,
	).Scan(&total)
	return total, err
}
