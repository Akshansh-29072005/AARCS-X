package logger

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

func NewLogger(env string, levelStr string) zerolog.Logger {

	zerolog.TimeFieldFormat = time.RFC3339

	level, err := zerolog.ParseLevel(strings.ToLower(levelStr))

	if err != nil {
		level = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(level)

	var logger zerolog.Logger

	if env == "release" {

		logger = zerolog.New(os.Stdout).
			With().
			Str("service", "AARCS-X").
			Timestamp().
			Logger()

	} else {

		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}

		logger = zerolog.New(consoleWriter).
			With().
			Str("service", "AARCS-X").
			Timestamp().
			Logger()
	}

	return logger
}