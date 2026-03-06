package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func GetLogger(c *gin.Context) zerolog.Logger {

	logger, exists := c.Get(LoggerKey)
	if !exists {
		return zerolog.Nop()
	}
	return logger.(zerolog.Logger)
}