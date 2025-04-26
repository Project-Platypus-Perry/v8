package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/project-platypus-perry/v8/internal/constants"
	"github.com/project-platypus-perry/v8/internal/model"
	"github.com/project-platypus-perry/v8/internal/service"
	"github.com/project-platypus-perry/v8/pkg/response"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
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
	var user model.User
	if err := c.Bind(&user); err != nil {
		return response.ValidationError(c, "Invalid request payload")
	}

	createdUser, err := h.userService.CreateUser(c.Request().Context(), &user)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusCreated, createdUser)
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
	if _, err := uuid.Parse(id); err != nil {
		return response.ValidationError(c, "Invalid user ID format")
	}

	user, err := h.userService.GetUser(c.Request().Context(), id)
	if err != nil {
		if err == constants.ErrNotFound {
			return response.NotFound(c, "User not found")
		}
		return response.InternalError(c)
	}

	return response.Success(c, http.StatusOK, user)
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
	if _, err := uuid.Parse(id); err != nil {
		return response.ValidationError(c, "Invalid user ID format")
	}

	var user model.User
	if err := c.Bind(&user); err != nil {
		return response.ValidationError(c, "Invalid request payload")
	}

	updatedUser, err := h.userService.UpdateUser(c.Request().Context(), id, &user)
	if err != nil {
		if err == constants.ErrNotFound {
			return response.NotFound(c, "User not found")
		}
		return response.InternalError(c)
	}

	return response.Success(c, http.StatusOK, updatedUser)
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
	if _, err := uuid.Parse(id); err != nil {
		return response.ValidationError(c, "Invalid user ID format")
	}

	if err := h.userService.DeleteUser(c.Request().Context(), id); err != nil {
		if err == constants.ErrNotFound {
			return response.NotFound(c, "User not found")
		}
		return response.InternalError(c)
	}

	return response.Success(c, http.StatusOK, nil)
}
