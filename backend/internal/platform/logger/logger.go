package logger

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

func NewLogger(env string, levelStr string) zerolog.Logger {

	// Setting time format for zerolog to RFC3339
	zerolog.TimeFieldFormat = time.RFC3339

	level, err := zerolog.ParseLevel(strings.ToLower(levelStr))

	if levelStr == "" {
		level = zerolog.InfoLevel
	}

	if err != nil {
		level = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(level)

	var logger zerolog.Logger

	if env == "release" {

		logger = zerolog.New(os.Stdout).
			With().
			Str("service", "aarcs-x").
			Timestamp().
			Int("pid", os.Getpid()).
			Caller().
			Logger()

	} else {

		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}

		logger = zerolog.New(consoleWriter).
			With().
			Str("service", "aarcs-x").
			Timestamp().
			Int("pid", os.Getpid()).
			Caller().
			Logger()
	}

	return logger
}