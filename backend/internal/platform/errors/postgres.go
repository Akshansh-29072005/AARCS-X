package errors

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func FromPostgresError(err error) error {

	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr){
		switch pgErr.Code {
		case "23505":
			return Conflict("user already exists", err)
		case "23503":
			return BadRequest("invalid reference", err)
		case "23502":
			return BadRequest("missing required field", err)
		}
	}
	return Internal("database error", err)
}