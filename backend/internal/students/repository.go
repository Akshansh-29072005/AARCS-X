package students

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateStudent(pool *pgxpool.Pool, FirstName string, LastName string, Email string, Phone string, Semester int, Branch string) (*Student, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
		INSERT INTO students (first_name, last_name, email, phone, semester, branch, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
		RETURNING id, first_name, last_name, email, phone, semester, branch, created_at, updated_at
	`

	var students Student
	var err error = pool.QueryRow(ctx, query, FirstName, LastName, Email, Phone, Semester, Branch).Scan(
		&students.ID,
		&students.FirstName,
		&students.LastName,
		&students.Email,
		&students.Phone,
		&students.Semester,
		&students.Branch,
		&students.CreatedAt,
		&students.UpdatedAt,
	)

	if err != nil{
		return nil, err
	}

	return &students, nil
}
func GetAllStudents(pool *pgxpool.Pool) ([]Student, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
		SELECT id, first_name, last_name, email, phone, semester, branch, created_at, updated_at
		FROM students
	`

	rows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student

	for rows.Next() {
		var student Student
		err := rows.Scan(
			&student.ID,
			&student.FirstName,
			&student.LastName,
			&student.Email,
			&student.Phone,
			&student.Semester,
			&student.Branch,
			&student.CreatedAt,
			&student.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}