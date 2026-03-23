package students

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/database"
	"github.com/redis/go-redis/v9"
)

const studentChacheTTL = time.Minute * 10

type Repository struct {
	db database.DBTX
	rdb *redis.Client
}

func NewRepository(db database.DBTX, rdb *redis.Client) *Repository {
	return &Repository{
		db: db,
		rdb: rdb,
	}
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

// CreateUser creates a user entry in the users table for authentication
func (r *Repository) CreateUser(ctx context.Context, email, hashedPassword string, studentID int) error {
	query := `
		INSERT INTO users (email, password, role, reference_id)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.Exec(ctx, query, email, hashedPassword, "student", studentID)
	return err
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

func (r *Repository) GetByID(ctx context.Context, id int) (*GetByIDStudentResponse, bool, error) {
	var student GetByIDStudentResponse

	// Check Redis cache first
	cacheKey := fmt.Sprintf("student:v1:%d", id)
	cachedData, err := r.rdb.Get(ctx, cacheKey).Bytes()
	if err == nil {
		// Cache hit, unmarshal cachedData into student
		if err := json.Unmarshal(cachedData, &student); err != nil {
			return nil, false, err
		}

		return &student, true, nil

	} else if err != redis.Nil {
		// Redis error 
		
	}

	// Cache miss, query PostgreSQL
	err = r.db.QueryRow(ctx,
		`SELECT id, name, email, phone, semester_id, department_id, institution_id FROM students WHERE id = $1`, id,
	).Scan(
		&student.ID,
		&student.Name,
		&student.Email,
		&student.Phone,
		&student.SemesterId,
		&student.DepartmentId,
		&student.InstitutionId,
	)
	
	if errors.Is(err, pgx.ErrNoRows){
		return nil, false, err
	}
	
	// Store result in Redis cache for future requests
	cachedData, err = json.Marshal(student)
	if err == nil {
		_ = r.rdb.Set(ctx, cacheKey, cachedData, studentChacheTTL).Err()
	}
	return &student, false, nil

}