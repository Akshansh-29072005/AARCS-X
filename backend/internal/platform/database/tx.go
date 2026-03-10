package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func WithTransaction(
	ctx context.Context,
	pool *pgxpool.Pool,
	fn func(tx DBTX) error,
) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)
	
	if err := fn(tx); err != nil {
		return err
	}

	return tx.Commit(ctx)
}