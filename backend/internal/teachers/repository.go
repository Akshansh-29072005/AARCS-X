package teachers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

const teacherChacheTTL = time.Minute * 10

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

// CreateUser creates a user entry in the users table for authentication
func (r *Repository) CreateUser(ctx context.Context, email, hashedPassword string, teacherID int) error {
	query := `
		INSERT INTO users (email, password, role, reference_id)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.Exec(ctx, query, email, hashedPassword, "teacher", teacherID)
	return err
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

func (r *Repository) GetByID (ctx context.Context, id int) (*GetByIDTeacherResponse, bool, error) {
	var teacher GetByIDTeacherResponse
	
	// Check Redis cache first
	cacheKey := fmt.Sprintf("teacher:v1:%d", id)
	cachedData, err := r.rdb.Get(ctx, cacheKey).Bytes()
	if err == nil {
		// Cache hit, unmarshal cachedData into institution
		if err := json.Unmarshal(cachedData, &teacher); err != nil {
			return nil, false, err
		}

		return &teacher, true, nil

	} else if err != redis.Nil {
		// Redis error 
		
	}

	// Cache miss, query PostgreSQL
	err = r.db.QueryRow(ctx,
		`SELECT id, name, email, phone, department_id, designation FROM teachers WHERE id = $1`, id,
	).Scan(
		&teacher.ID,
		&teacher.Name,
		&teacher.Email,
		&teacher.Phone,
		&teacher.DepartmentId,
		&teacher.Designation,
	)
	
	if errors.Is(err, pgx.ErrNoRows){
		return nil, false, err
	}
	
	// Store result in Redis cache for future requests
	cachedData, err = json.Marshal(teacher)
	if err == nil {
		_ = r.rdb.Set(ctx, cacheKey, cachedData, teacherChacheTTL).Err()
	}
	return &teacher, false, nil
}