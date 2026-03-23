package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func WithTransaction(
	ctx context.Context,
	pool *pgxpool.Pool,
	fn func(tx DBTX, rdb *redis.Client) error,
) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)
	
	if err := fn(tx, nil); err != nil {
		return err
	}

	return tx.Commit(ctx)
}