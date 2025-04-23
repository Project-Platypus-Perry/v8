package handler

import (
	"net/http"

	"github.com/gagan-gaurav/v8/internal/model"
	"github.com/gagan-gaurav/v8/internal/service"
	"github.com/gagan-gaurav/v8/pkg/logger"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// @Summary Create a new user
// @Description Creates a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "User object that needs to be created"
// @Success 201 {object} model.User "User created successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users [post]
func (h *UserHandler) CreateUser(c echo.Context) error {
	req := new(model.User)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Generate a new UUID for the user
	req.ID = uuid.New().String()

	logger.Info("Creating user", zap.Any("user", req))

	createdUser, err := h.service.CreateUser(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdUser)
}

// @Summary Get user by ID
// @Description Retrieves a user by their UUID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User UUID" format(uuid)
// @Success 200 {object} model.User "User found successfully"
// @Failure 400 {object} map[string]string "Invalid UUID format"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")

	// validate uuid
	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	user, err := h.service.GetUser(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

// @Summary Update user
// @Description Updates an existing user's information
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User UUID" format(uuid)
// @Param user body model.User true "Updated user object"
// @Success 200 {object} model.User "User updated successfully"
// @Failure 400 {object} map[string]string "Invalid UUID format or request body"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users/{id} [patch]
func (h *UserHandler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	user := new(model.User)

	user.Name = c.FormValue("name")
	user.Password = c.FormValue("password")

	logger.Info("Updating user", zap.Any("user", user), zap.String("id", id), zap.String("name", user.Name), zap.String("password", user.Password))
	// validate uuid
	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updatedUser, err := h.service.UpdateUser(c.Request().Context(), id, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedUser)
}

// @Summary Delete user
// @Description Deletes a user by their UUID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User UUID" format(uuid)
// @Success 200 {string} string "User deleted successfully"
// @Failure 400 {object} map[string]string "Invalid UUID format"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	// validate uuid
	if _, err := uuid.Parse(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	err := h.service.DeleteUser(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "User deleted successfully")
}
