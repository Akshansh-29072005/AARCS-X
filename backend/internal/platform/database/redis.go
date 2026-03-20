package database

import (
	"context"

	"github.com/Akshansh-29072005/AARCS-X/backend/internal/config"
	"github.com/Akshansh-29072005/AARCS-X/backend/internal/platform/logger"
	"github.com/redis/go-redis/v9"
)

func ConnectRedis(redisURL string) (*redis.Client, error) {

	var cfg *config.Config
	var err error
	cfg, err = config.Load()
	if err != nil {
		panic("failed to load configuration: " + err.Error())
	}

	appLogger := logger.NewLogger(cfg.GinMode, cfg.LogLevel)

	var ctx context.Context = context.Background()

	config, err := redis.ParseURL(redisURL)
	
	if err != nil {
		appLogger.Fatal().Err(err).Msg("Unable to parse Redis URL")
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr: config.Addr,
		DB: 0,
	})
	
	err = client.Ping(ctx).Err()

	if err != nil {
		appLogger.Fatal().Err(err).Msg("Unable to connect to Redis")
		return nil, err
	}

	appLogger.Info().Msg("Successfully connected to Redis")
	return client, nil
}