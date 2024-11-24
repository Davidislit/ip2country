package middleware

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRateLimiterAllow(t *testing.T) {
	rateLimiter := NewRateLimiter(2)

	assert.True(t, rateLimiter.Allow())

	assert.True(t, rateLimiter.Allow())

	assert.False(t, rateLimiter.Allow())

	time.Sleep(1 * time.Second)

	assert.True(t, rateLimiter.Allow())
}
