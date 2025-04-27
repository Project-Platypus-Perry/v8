package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/project-platypus-perry/v8/internal/model"
	"github.com/project-platypus-perry/v8/internal/service"
	"github.com/project-platypus-perry/v8/pkg/logger"
	"github.com/project-platypus-perry/v8/pkg/response"
	"go.uber.org/zap"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	loginRequest := struct {
		Email    string `json:"Email" validate:"required,email"`
		Password string `json:"Password" validate:"required"`
	}{}

	if err := c.Bind(&loginRequest); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request format: "+err.Error())
	}

	// validate email and password
	if err := c.Validate(&loginRequest); err != nil {
		return response.Error(c, http.StatusBadRequest, "Validation error: "+err.Error())
	}

	logger.Info("Logging in", zap.Any("loginRequest", loginRequest))

	user, tokenPair, err := h.authService.Login(c.Request().Context(), loginRequest.Email, loginRequest.Password)
	if err != nil {
		return response.Error(c, http.StatusUnauthorized, err.Error())
	}

	return response.Success(c, http.StatusOK, map[string]interface{}{
		"User":      user,
		"TokenPair": tokenPair,
	})
}

// Register an Organization with an Admin User.
func (h *AuthHandler) RegisterOrganization(c echo.Context) error {
	userData := model.User{}

	if err := c.Bind(&userData); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request format: "+err.Error())
	}

	if err := c.Validate(&userData); err != nil {
		return response.Error(c, http.StatusBadRequest, "Validation error: "+err.Error())
	}

	err := h.authService.RegisterOrganization(
		c.Request().Context(),
		&userData,
		&userData.Organization,
	)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusCreated, "Organization created and Admin registered successfully")
}
