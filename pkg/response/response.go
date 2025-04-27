package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Response represents the standard API response structure
type Response struct {
	Success bool        `json:"Success"`
	Data    interface{} `json:"Data"`
	Message string      `json:"Message"`
	Code    int         `json:"StatusCode"`
}

// Success sends a successful response
func Success(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, Response{
		Success: true,
		Data:    data,
		Message: "",
		Code:    code,
	})
}

// Error sends an error response
func Error(c echo.Context, code int, message string) error {
	return c.JSON(code, Response{
		Success: false,
		Data:    nil,
		Message: message,
		Code:    code,
	})
}

// ValidationError sends a validation error response
func ValidationError(c echo.Context, message string) error {
	return Error(c, http.StatusBadRequest, message)
}

// NotFound sends a not found error response
func NotFound(c echo.Context, message string) error {
	return Error(c, http.StatusNotFound, message)
}

// InternalError sends an internal server error response
func InternalError(c echo.Context) error {
	return Error(c, http.StatusInternalServerError, "Internal Server Error")
}
