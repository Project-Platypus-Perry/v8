package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// @Summary Health check endpoint
// @Description Check if the API service is up and running
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {string} string "OK"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /health [get]
func (h *HealthHandler) Check(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
