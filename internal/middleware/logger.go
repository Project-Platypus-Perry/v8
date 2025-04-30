package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/project-platypus-perry/v8/pkg/logger"
	"go.uber.org/zap"
)

func RequestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Log request details
		start := time.Now()
		req := c.Request()
		// logger.Info("Request",
		// 	zap.String("method", req.Method),
		// 	zap.String("path", req.URL.Path),
		// 	zap.String("remote_ip", c.RealIP()),
		// 	zap.String("user_agent", req.UserAgent()),
		// )

		// Process request
		err := next(c)

		// Log response details
		latency := time.Since(start)
		logger.Info("Response",
			zap.String("method", req.Method),
			zap.String("path", req.URL.Path),
			zap.Int("status", c.Response().Status),
			zap.Duration("latency", latency),
		)

		return err
	}
}
