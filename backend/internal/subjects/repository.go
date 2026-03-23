package subjects

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

const subjectChacheTTL = time.Minute * 10

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

func (r *Repository) GetByID(ctx context.Context, id int) (*GetByIDSubjectResponse, bool, error) {
	var subject GetByIDSubjectResponse

	// Check Redis cache first
	cacheKey := fmt.Sprintf("subject:v1:%d", id)
	cachedData, err := r.rdb.Get(ctx, cacheKey).Bytes()
	if err == nil {
		// Cache hit, unmarshal cachedData into institution
		if err := json.Unmarshal(cachedData, &subject); err != nil {
			return nil, false, err
		}

		return &subject, true, nil

	} else if err != redis.Nil {
		// Redis error 
		
	}

	// Cache miss, query PostgreSQL
	err = r.db.QueryRow(ctx,
		`SELECT id, name, code FROM subjects WHERE id = $1`, id,
	).Scan(
		&subject.ID,
		&subject.Name,
		&subject.Code,
	)
	
	if errors.Is(err, pgx.ErrNoRows){
		return nil, false, err
	}
	
	// Store result in Redis cache for future requests
	cachedData, err = json.Marshal(subject)
	if err == nil {
		_ = r.rdb.Set(ctx, cacheKey, cachedData, subjectChacheTTL).Err()
	}
	return &subject, false, nil

}