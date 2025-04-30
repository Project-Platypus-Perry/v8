package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/project-platypus-perry/v8/internal/config"
)

type visitor struct {
	requests  int
	startTime time.Time
}

type RateLimiter struct {
	visitors map[string]*visitor
	mu       sync.RWMutex
	config   *config.RateLimiterConfig
}

func NewRateLimiter(cfg *config.RateLimiterConfig) *RateLimiter {
	return &RateLimiter{
		visitors: make(map[string]*visitor),
		config:   cfg,
	}
}

func (rl *RateLimiter) RateLimit(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ip := c.RealIP()

		rl.mu.Lock()
		defer rl.mu.Unlock()

		v, exists := rl.visitors[ip]
		now := time.Now()

		if !exists {
			// First request from this IP
			rl.visitors[ip] = &visitor{
				requests:  1,
				startTime: now,
			}
			return next(c)
		}

		// Check if the time window has expired
		if now.Sub(v.startTime) > rl.config.Duration {
			// Reset for new time window
			v.requests = 1
			v.startTime = now
			return next(c)
		}

		// If we're still within the time window
		if v.requests >= rl.config.Requests {
			return c.JSON(http.StatusTooManyRequests, map[string]string{
				"error": "Rate limit exceeded. Please try again later.",
			})
		}

		// Increment request count and allow
		v.requests++
		return next(c)
	}
}
