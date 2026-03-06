package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func LoggerMiddleware(baseLogger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		requestID := c.GetString(RequestIDKey)

		requestLogger := baseLogger.With().
			Str("request_id", requestID).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Logger()

		c.Set(LoggerKey, requestLogger)

		c.Next()
	}
}