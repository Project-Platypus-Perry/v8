package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/project-platypus-perry/v8/pkg/response"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthResponse struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

// @Summary Health check endpoint
// @Description Check if the API service is up and running
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=HealthResponse} "Service is healthy"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /health [get]
func (h *HealthHandler) Check(c echo.Context) error {
	return response.Success(c, http.StatusOK, HealthResponse{
		Status: "healthy",
		Time:   time.Now().Format(time.RFC3339),
	})
}
