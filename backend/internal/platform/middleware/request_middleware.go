package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// Skip infrastructure endpoints
		if path == "/api/v1/system/metrics" || path == "/api/v1/system/health" {
			c.Next()
			return
		}

		log := GetLogger(c)
		log.Info().
			Str("method", c.Request.Method).
			Str("path", path).
			Msg("Request started")

		start := time.Now()
		clientIP := c.ClientIP()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		log.Info().
			Str("client_ip", clientIP).
			Int("status", status).
			Dur("latency", latency).
			Msg("Request completed")
	}
}
