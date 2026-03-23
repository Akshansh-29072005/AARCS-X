package semesters

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

const semesterChacheTTL = time.Minute * 10

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

func (r *Repository) GetByID(ctx context.Context, id int) (*GetByIDSemesterResponse, bool, error) {
	var semester GetByIDSemesterResponse

	// Check Redis cache first
	cacheKey := fmt.Sprintf("semester:v1:%d", id)
	cachedData, err := r.rdb.Get(ctx, cacheKey).Bytes()
	if err == nil {
		// Cache hit, unmarshal cachedData into institution
		if err := json.Unmarshal(cachedData, &semester); err != nil {
			return nil, false, err
		}

		return &semester, true, nil

	} else if err != redis.Nil {
		// Redis error 
		
	}

	// Cache miss, query PostgreSQL
	err = r.db.QueryRow(ctx,
		`SELECT id, number, department_id FROM semesters WHERE id = $1`, id,
	).Scan(
		&semester.ID,
		&semester.Number,
		&semester.DepartmentId,
	)
	
	if errors.Is(err, pgx.ErrNoRows){
		return nil, false, err
	}
	
	// Store result in Redis cache for future requests
	cachedData, err = json.Marshal(semester)
	if err == nil {
		_ = r.rdb.Set(ctx, cacheKey, cachedData, semesterChacheTTL).Err()
	}
	return &semester, false, nil

}