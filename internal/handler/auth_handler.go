package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/project-platypus-perry/v8/internal/model"
	"github.com/project-platypus-perry/v8/internal/service"
	"github.com/project-platypus-perry/v8/pkg/jwt"
	"github.com/project-platypus-perry/v8/pkg/response"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// LoginResponse represents the successful login response
type LoginResponse struct {
	User      *model.User    `json:"user"`
	TokenPair *jwt.TokenPair `json:"tokenPair"`
}

// @Summary User login
// @Description Authenticate user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param login body model.LoginRequest true "Login credentials"
// @Success 200 {object} response.Response{data=LoginResponse} "Login successful"
// @Failure 400 {object} response.Response "Invalid request format or validation error"
// @Failure 401 {object} response.Response "Invalid credentials"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	loginRequest := model.LoginRequest{}

	if err := c.Bind(&loginRequest); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request format: "+err.Error())
	}

	// validate email and password
	if err := c.Validate(&loginRequest); err != nil {
		return response.Error(c, http.StatusBadRequest, "Validation error: "+err.Error())
	}

	user, tokenPair, err := h.authService.Login(c.Request().Context(), loginRequest.Email, loginRequest.Password)
	if err != nil {
		return response.Error(c, http.StatusUnauthorized, err.Error())
	}

	return response.Success(c, http.StatusOK, LoginResponse{
		User:      user,
		TokenPair: tokenPair,
	})
}

// @Summary Register organization
// @Description Register a new organization with an admin user
// @Tags auth
// @Accept json
// @Produce json
// @Param registration body model.User true "Organization and admin user details"
// @Success 201 {object} response.Response "Organization created and admin registered successfully"
// @Failure 400 {object} response.Response "Invalid request format or validation error"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /auth/register [post]
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
