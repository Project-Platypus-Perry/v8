package middleware

import (
	"github.com/labstack/echo/v4"
)

func RateLimit(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Rate limit logic
		return next(c)
	}
}
