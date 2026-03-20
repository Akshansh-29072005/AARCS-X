package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GinMode	     string
	DatabaseURL  string
	RedisURL	 string
	Port         string
	JWTSecret    string
	LogLevel     string
}

func Load()(*Config, error) {
	var err error =godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	var config*Config = &Config{
		GinMode:     os.Getenv("GIN_MODE"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		RedisURL:    os.Getenv("REDIS_URL"),
		Port:        os.Getenv("PORT"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		LogLevel:    os.Getenv("LOG_LEVEL"),
	}

	return config, nil
}