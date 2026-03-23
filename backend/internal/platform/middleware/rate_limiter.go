package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type RateLimiter struct {
	// Implementation of rate limiter
	rdb 	   *redis.Client
	limit 	   int64
	windowSize time.Duration
}

func NewRateLimiter(rdb *redis.Client, limit int64, window time.Duration) *RateLimiter {
	return &RateLimiter{
		rdb:        rdb,
		limit:      limit,
		windowSize: window,
	}
}

func (rl *RateLimiter) RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implementation of rate limiting logic
		ctx := c.Request.Context()
		log := GetLogger(c)

		// Get client IP address
		clientIP := c.ClientIP()
		key := "rate_limit:v1:" + clientIP

		// Increment the count for this IP
		count, err := rl.rdb.Incr(ctx, key).Result()
		if err != nil {
			log.Error().Err(err).Msg("redis unavailable, skipping rate limit")
			c.Next()
			return
		}

		if count == 1 {
			// Set expiration for the key if it's the first request
			_ = rl.rdb.Expire(ctx, key, rl.windowSize).Err()
		}

		if count > rl.limit{
			log.Warn().
				Str("client_ip", clientIP).
				Int64("count", count).
				Msg("rate limit exceeded")

			ttl, err := rl.rdb.TTL(ctx, key).Result()
			if err != nil {
				log.Error().Err(err).Msg("failed to get TTL for rate limit key")
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
					"error": "rate limit exceeded",
				})
				return
			}

			c.Header("Retry-After", fmt.Sprintf("%.0f", ttl.Seconds()))

			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
				"retry_after_seconds": int(ttl.Seconds()),
			})
			return
		}
		c.Next()
	}
}