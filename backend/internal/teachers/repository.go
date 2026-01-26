package teachers

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

func (r *Repository) Create(ctx context.Context, e *TeacherEntity) (*TeacherEntity, error) {
	query := `
        INSERT INTO teachers(name, email, phone, password, department_id, designation, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, NOW())
        RETURNING id, created_at`

	err := r.db.QueryRow(ctx, query,
		e.Name,
		e.Email,
		e.Phone,
		e.Password,
		e.DepartmentId,
		e.Designation,
	).Scan(&e.ID, &e.CreatedAt)

	return e, err
}

func (r *Repository) List(ctx context.Context, q GetTeachersRequest) ([]Teacher, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, name, department_id, designation
         FROM teachers
         WHERE ($1 = 0 OR department = $1)
         AND ($2 = '' OR designation = $2)`,
		q.DepartmentId, q.Designation,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var teachers []Teacher
	for rows.Next() {
		var t Teacher
		err := rows.Scan(
			&t.ID,
			&t.Name,
			&t.DepartmentId,
			&t.Designation,
		)
		if err != nil {
			return nil, err
		}

		teachers = append(teachers, t)
	}

	return teachers, nil
}

func (r *Repository) Count(ctx context.Context, q GetTeachersRequest) (int, error) {
	var total int
	err := r.db.QueryRow(ctx,
		`SELECT COUNT(*) FROM teachers
        WHERE ($1 = 0 OR department = $1)
        AND ($2 = '' OR designation = $2)`,
		q.DepartmentId, q.Designation,
	).Scan(&total)
	return total, err
}
