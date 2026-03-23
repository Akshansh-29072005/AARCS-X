package departments

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

const departmentChacheTTL = time.Minute * 10

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

func (r *Repository) Create(ctx context.Context, e *DepartmentEntity) (*DepartmentEntity, error) {
	query := `
		INSERT INTO departments (name, code, head_of_department, institution_id, created_at)
		VALUES ($1,$2,$3,$4,NOW())
		RETURNING id, created_at
	`

	err := r.db.QueryRow(ctx, query,
		e.Name,
		e.Code,
		e.HeadOfDepartment,
		e.InstitutionId,
	).Scan(&e.ID, &e.CreatedAt)

	return e, err
}

func (r *Repository) List(ctx context.Context, q GetDepartmentRequest) ([]Department, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, name, code, head_of_department, institution_id
		FROM departments
		WHERE ($1 = '' OR name = $1)
		AND ($2 = '' OR code = $2)
		AND ($3 = '' OR head_of_department = $3)
		AND ($4 = 0 OR institution_id = $4)`,
		q.Name, q.Code, q.HeadOfDepartment, q.InstitutionId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []Department
	for rows.Next() {
		var s Department
		err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.Code,
			&s.HeadOfDepartment,
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

func (r *Repository) Count(ctx context.Context, q GetDepartmentRequest) (int, error) {
	var total int
	err := r.db.QueryRow(ctx,
		`SELECT COUNT(*) FROM departments
			 WHERE ($1 = '' OR name = $1)
			 AND ($2 = '' OR code = $2)
			 AND ($3 = '' OR head_of_department = $3)
			 AND ($4 = 0 OR institution_id = $4)`,
		q.Name, q.Code, q.HeadOfDepartment, q.InstitutionId,
	).Scan(&total)
	return total, err
}

func (r *Repository) GetByID(ctx context.Context, id int) (*GetByIDDepartmentResponse, bool, error) {
	var department GetByIDDepartmentResponse

	// Check Redis cache first
	cacheKey := fmt.Sprintf("department:v1:%d", id)
	cachedData, err := r.rdb.Get(ctx, cacheKey).Bytes()
	if err == nil {
		// Cache hit, unmarshal cachedData into institution
		if err := json.Unmarshal(cachedData, &department); err != nil {
			return nil, false, err
		}

		return &department, true, nil

	} else if err != redis.Nil {
		// Redis error 
		
	}

	// Cache miss, query PostgreSQL
	err = r.db.QueryRow(ctx,
		`SELECT id, name, code FROM departments WHERE id = $1`, id,
	).Scan(
		&department.ID,
		&department.Name,
		&department.Code,
	)
	
	if errors.Is(err, pgx.ErrNoRows){
		return nil, false, err
	}
	
	// Store result in Redis cache for future requests
	cachedData, err = json.Marshal(department)
	if err == nil {
		_ = r.rdb.Set(ctx, cacheKey, cachedData, departmentChacheTTL).Err()
	}
	return &department, false, nil

}