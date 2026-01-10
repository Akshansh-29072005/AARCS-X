package teachers

import (
	"context"
	"time"

	"github.com/Akshansh-29072005/AARCS-X/backend/internals/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTeacher(pool *pgxpool.Pool, FirstName string, LastName string, Email string, Phone string, Department string, Designation string) (*models.Teacher, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
		INSERT INTO teachers (first_name, last_name, email, phone, department, designation, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
		RETURNING id, first_name, last_name, email, phone, department, designation, created_at, updated_at
	`

	var teachers models.Teacher
	var err error = pool.QueryRow(ctx, query, FirstName, LastName, Email, Phone, Department, Designation).Scan(
		&teachers.ID,
		&teachers.FirstName,
		&teachers.LastName,
		&teachers.Email,
		&teachers.Phone,
		&teachers.Department,
		&teachers.Designation,
		&teachers.CreatedAt,
		&teachers.UpdatedAt,
	)

	if err != nil{
		return nil, err
	}

	return &teachers, nil
}