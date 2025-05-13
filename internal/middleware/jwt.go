package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/project-platypus-perry/v8/internal/config"
	"github.com/project-platypus-perry/v8/pkg/jwt"
	"github.com/project-platypus-perry/v8/pkg/response"
)

type JWTMiddleware struct {
	config *config.JWTConfig
}

func NewJWTMiddleware(cfg *config.JWTConfig) *JWTMiddleware {
	return &JWTMiddleware{
		config: cfg,
	}
}

func (m *JWTMiddleware) JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return response.Error(c, http.StatusUnauthorized, "No authorization header")
		}

		// Check if the Authorization header has the correct format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return response.Error(c, http.StatusUnauthorized, "Invalid authorization header format")
		}

		tokenString := parts[1]
		claims, err := jwt.ValidateToken(tokenString, m.config.AccessTokenSecret, jwt.AccessToken)
		if err != nil {
			return response.Error(c, http.StatusUnauthorized, "Invalid or expired token")
		}

		// Store user information in context
		c.Set("UserID", claims.UserID)
		c.Set("Role", claims.Role)
		c.Set("OrganizationID", claims.OrganizationID)
		return next(c)
	}
}

func (m *JWTMiddleware) RefreshToken(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return response.Error(c, http.StatusUnauthorized, "No authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return response.Error(c, http.StatusUnauthorized, "Invalid authorization header format")
	}

	refreshTokenString := parts[1]
	claims, err := jwt.ValidateToken(refreshTokenString, m.config.RefreshTokenSecret, jwt.RefreshToken)
	if err != nil {
		return response.Error(c, http.StatusUnauthorized, err.Error())
	}

	// Generate new token pair
	tokenPair, err := jwt.GenerateTokenPair(claims.UserID, claims.Role, claims.OrganizationID, m.config)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, tokenPair)
}
