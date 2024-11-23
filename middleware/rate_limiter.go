package middleware

import (
	"ip2country/api"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiter struct {
	mutex     sync.Mutex
	timestamp time.Time
	count     int
	limit     int
	interval  time.Duration
}

func NewRateLimiter(limit int) *RateLimiter {
	return &RateLimiter{
		timestamp: time.Now(),
		limit:     limit,
		interval:  time.Second,
	}
}

func (rateLimiter *RateLimiter) Allow() bool {
	rateLimiter.mutex.Lock()
	defer rateLimiter.mutex.Unlock()

	now := time.Now()
	if now.Sub(rateLimiter.timestamp) > rateLimiter.interval {
		rateLimiter.timestamp = now
		rateLimiter.count = 0
	}

	if rateLimiter.count < rateLimiter.limit {
		rateLimiter.count++
		return true
	}

	return false
}

func RateLimitMiddleware(rateLimiter *RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if rateLimiter.Allow() {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, api.ResponseError{Error: "rate limit exceeded"})
			return
		}
	}
}
