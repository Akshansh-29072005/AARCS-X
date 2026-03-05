package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func RequestLogger(log zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		path := c.Request.URL.Path
		method := c.Request.Method
		clientIP := c.ClientIP()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		log.Info().
			Str("method", method).
			Str("path", path).
			Str("client_ip", clientIP).
			Int("status", status).
			Dur("latency", latency).
			Msg("Request completed")
	}
}