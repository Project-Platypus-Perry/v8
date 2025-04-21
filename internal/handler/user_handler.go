package handler

import (
	"net/http"

	"github.com/gagan-gaurav/base/internal/service"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) ListUsers(c echo.Context) error {
	users, _ := h.service.ListUsers(c.Request().Context())
	return c.JSON(http.StatusOK, users)
}
