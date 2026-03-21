package institutes

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

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

func (r *Repository) Create(ctx context.Context, e *InstitutionEntity) (*InstitutionEntity, error) {
	query := `
		INSERT INTO institutions (name, code, password, created_at)
		VALUES ($1,$2,$3,NOW())
		RETURNING id, created_at
	`

	err := r.db.QueryRow(ctx, query,
		e.Name,
		e.Code,
		e.Password,
	).Scan(&e.ID, &e.CreatedAt)

	return e, err
}

func (r *Repository) List(ctx context.Context, q GetInstitutionRequest) ([]Institute, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, name, code
		FROM institutions
		WHERE ($1 = '' OR name = $1)
		AND ($2 = '' OR code = $2)`,
		q.Name, q.Code,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var institutions []Institute
	for rows.Next() {
		var s Institute
		err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.Code,
		)
		if err != nil {
			return nil, err
		}
		institutions = append(institutions, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return institutions, nil
}

func (r *Repository) Count(ctx context.Context, q GetInstitutionRequest) (int, error) {
	var total int
	err := r.db.QueryRow(ctx,
		`SELECT COUNT(*) FROM institutions
			 WHERE ($1 = '' OR name = $1)
			 AND ($2 = '' OR code = $2)`,
		q.Name, q.Code,
	).Scan(&total)
	return total, err
}

func (r *Repository) GetByID(ctx context.Context, id int) (*GetByIDInstituteRequest, error) {
	var institution GetByIDInstituteRequest

	// Check Redis cache first
	cacheKey := fmt.Sprintf("institution:%d", id)
	cachedData, err := r.rdb.Get(ctx, cacheKey).Bytes()
	if err == nil {
		// Cache hit, unmarshal cachedData into institution
		json.Unmarshal(cachedData, &institution)
		return &institution, nil

	}
	// Cache miss, query the database
	err = r.db.QueryRow(ctx,
		`SELECT id, name, code FROM institutions WHERE id = $1`, id,
	).Scan(&institution.ID, &institution.Name, &institution.Code)

	if err != nil {
		return nil, err
	}

	// Store the result in Redis cache for future requests
	data,_ := json.Marshal(institution)
	r.rdb.Set(ctx, cacheKey, data, time.Minute*5) // Cache for 5 minutes

	return &institution, nil
}