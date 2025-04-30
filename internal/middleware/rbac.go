package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/project-platypus-perry/v8/internal/constants"
)

// RequirePermission middleware checks if the user has the required permission
func RequirePermission(requiredPermission constants.Permission) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get user role from context (set by JWT middleware)
			role, ok := c.Get("user_role").(constants.UserRole)
			if !ok {
				return c.JSON(http.StatusForbidden, map[string]string{
					"error": "Role not found in token",
				})
			}

			// Check if role has the required permission
			permissions := constants.RolePermissions[role]
			hasPermission := false
			for _, permission := range permissions {
				if permission == requiredPermission {
					hasPermission = true
					break
				}
			}

			if !hasPermission {
				return c.JSON(http.StatusForbidden, map[string]string{
					"error": "Insufficient permissions",
				})
			}

			return next(c)
		}
	}
}

// RequireRole middleware checks if the user has the required role
func RequireRole(requiredRoles ...constants.UserRole) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get user role from context (set by JWT middleware)
			role, ok := c.Get("user_role").(constants.UserRole)
			if !ok {
				return c.JSON(http.StatusForbidden, map[string]string{
					"error": "Role not found in token",
				})
			}

			// Check if user's role matches any of the required roles
			hasRole := false
			for _, r := range requiredRoles {
				if role == r {
					hasRole = true
					break
				}
			}

			if !hasRole {
				return c.JSON(http.StatusForbidden, map[string]string{
					"error": "Insufficient role",
				})
			}

			return next(c)
		}
	}
}
