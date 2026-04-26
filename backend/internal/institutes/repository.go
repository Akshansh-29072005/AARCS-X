package institutes

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

const institutionChacheTTL = 10 * time.Minute
const RoleInstitutionOwner string = "institution_owner"
const RoleInstitutionAdmin string = "institution_admin"

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

func (r *Repository) Create(ctx context.Context, e *InstitutionEntity, userID int, roleName string) (*InstitutionEntity, error) {

	roleName = RoleInstitutionOwner
	query := `
	WITH inserted_institution AS (
		INSERT INTO institutions (name, code, official_email, address, district, state, country, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
		RETURNING id, created_at
	),
	inserted_owner AS (
		INSERT INTO institution_owners (institution_id, user_id)
		SELECT id, $8 FROM inserted_institution
		RETURNING institution_id
	),
	inserted_role AS (
		INSERT INTO roles (user_id, role, created_at)
		VALUES ($8, $9, NOW())
		RETURNING id, created_at
	)
	SELECT 
		id,
		created_at
	FROM inserted_institution;
	`

	err := r.db.QueryRow(ctx, query,
		e.Name,
		e.Code,
		e.OfficialEmail,
		e.Address,
		e.District,
		e.State,
		e.Country,
		userID,
		roleName,
	).Scan(&e.ID, &e.CreatedAt)

	if err != nil {
		return nil, err
	}

	return e, err
}

func (r *Repository) IsInstitutionOwner(ctx context.Context, institutionID int, userID int) (bool, error) {
	var exists bool
	err := r.db.QueryRow(ctx,
		`SELECT EXISTS (
			SELECT 1 FROM institution_owners 
			WHERE institution_id = $1 AND user_id = $2
		)`, institutionID, userID,
	).Scan(&exists)

	return exists, err
}

func (r *Repository) PromoteToAdmin(ctx context.Context, institutionID int, userID int) (*AdminEntity, error) {	
	query := `
	WITH new_role AS (
		INSERT INTO roles (user_id, role, created_at)
		VALUES ($1, $2, NOW())
		ON CONFLICT (user_id, role) DO UPDATE SET created_at = NOW()
		RETURNING id, role, created_at
	),
	inserted_admin AS (
		INSERT INTO institution_admins (user_id, institution_id, created_at)
		VALUES ($1, $3, NOW())
		ON CONFLICT (user_id, institution_id) DO NOTHING
		RETURNING instituion_id
	)
	SELECT
		    nr.user_id,
		    nr.role,
		    nr.created_at,
			$3 AS institution_id
	FROM new_role nr
	`

	var e AdminEntity

	err := r.db.QueryRow(ctx,
		query,
		userID,
		RoleInstitutionAdmin,
		institutionID,
		).Scan(
			&e.UserID,
			&e.Role,
			&e.CreatedAt,
			&e.InstitutionID,
		)
	if err != nil {
		return nil, err
	}

	return &e, nil
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

func (r *Repository) GetByID(ctx context.Context, id int) (*GetByIDInstituteResponse,bool, error) {
	var institution GetByIDInstituteResponse

	// Check Redis cache first
	cacheKey := fmt.Sprintf("institution:v1:%d", id)
	cachedData, err := r.rdb.Get(ctx, cacheKey).Bytes()
	if err == nil {
		// Cache hit, unmarshal cachedData into institution
		if err := json.Unmarshal(cachedData, &institution); err != nil {
			return nil, false, err
		}

		return &institution, true, nil

	} else if err != redis.Nil {
		// Redis error
		return nil, false, err
	}

	// Cache miss, query PostgreSQL
	err = r.db.QueryRow(ctx,
		`SELECT id, name, code FROM institutions WHERE id = $1`, id,
	).Scan(
		&institution.ID,
		&institution.Name,
		&institution.Code,
	)
	
	if errors.Is(err, pgx.ErrNoRows){
		return nil, false, err
	}
	
	// Store result in Redis cache for future requests
	cachedData, err = json.Marshal(institution)
	if err == nil {
		_ = r.rdb.Set(ctx, cacheKey, cachedData, institutionChacheTTL).Err()
	}
	return &institution, false, nil

}