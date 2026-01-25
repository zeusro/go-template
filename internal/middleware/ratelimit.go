package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"zeusro.com/hermes/internal/core/config"
)

// RateLimiter provides rate limiting middleware
type RateLimiter struct {
	limiter *rate.Limiter
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(cfg config.Config) *RateLimiter {
	requests := cfg.RateLimit.Requests
	if requests == 0 {
		requests = 100 // default: 100 requests per second
	}
	burst := cfg.RateLimit.Burst
	if burst == 0 {
		burst = requests
	}

	limiter := rate.NewLimiter(rate.Limit(requests), burst)

	return &RateLimiter{
		limiter: limiter,
	}
}

// Middleware returns a Gin middleware for rate limiting
func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !rl.limiter.Allow() {
			c.Header("X-RateLimit-Limit", strconv.Itoa(rl.limiter.Burst()))
			c.Header("X-RateLimit-Remaining", "0")
			c.JSON(http.StatusTooManyRequests, gin.H{
				"code":    http.StatusTooManyRequests,
				"message": "Rate limit exceeded",
			})
			c.Abort()
			return
		}

		remaining := rl.limiter.Burst() - 1 // Simplified
		c.Header("X-RateLimit-Limit", strconv.Itoa(rl.limiter.Burst()))
		c.Header("X-RateLimit-Remaining", strconv.Itoa(remaining))

		c.Next()
	}
}
