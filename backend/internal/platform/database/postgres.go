package database

import (
	"context"

	"github.com/Akshansh-29072005/AARCS-X/backend/internal/config"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseURL string) (*pgxpool.Pool, error) {

	var cfg *config.Config
	var err error
	cfg, err = config.Load()
	if err != nil {
		panic("failed to load configuration: " + err.Error())
	}
	appLogger := logger.NewLogger(cfg.GinMode, cfg.LogLevel)

	var ctx context.Context = context.Background()

	var config *pgxpool.Config
	config, err = pgxpool.ParseConfig(databaseURL)

	if err != nil {
		appLogger.Fatal().Err(err).Msg("Unable to parse database URL")
		return nil, err
	}

	var pool *pgxpool.Pool
	pool, err = pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		appLogger.Fatal().Err(err).Msg("Unable to create connection pool")
		return nil, err 
	}

	err = pool.Ping(ctx)

	if err != nil {
		appLogger.Fatal().Err(err).Msg("Unable to connect to database")
		pool.Close()
		return nil, err
	}

	appLogger.Info().Msg("Connected to the database successfully")
	return pool, nil
}