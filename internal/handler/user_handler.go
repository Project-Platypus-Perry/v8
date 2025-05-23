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

// // @Summary Create a new user
// // @Description Creates a new user with the provided information
// // @Tags users
// // @Accept json
// // @Produce json
// // @Param user body model.User true "User object that needs to be created"
// // @Success 201 {object} model.User "User created successfully"
// // @Failure 400 {object} map[string]string "Invalid request body"
// // @Failure 500 {object} map[string]string "Internal server error"
// // @Router /users [post]
// func (h *UserHandler) CreateUser(c echo.Context) error {
// 	var user model.User
// 	if err := c.Bind(&user); err != nil {
// 		return response.ValidationError(c, "Invalid request payload")
// 	}

// 	createdUser, err := h.userService.CreateUser(c.Request().Context(), &user)
// 	if err != nil {
// 		return response.Error(c, http.StatusInternalServerError, err.Error())
// 	}

// 	return response.Success(c, http.StatusCreated, createdUser)
// }

// @Summary Get user by ID or email
// @Description Retrieves a user by their UUID or email address
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query string false "User UUID" format(uuid)
// @Param email query string false "User email address" format(email)
// @Success 200 {object} response.Response{data=model.User} "User found successfully"
// @Failure 400 {object} response.Response "Invalid parameters or validation error"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 404 {object} response.Response "User not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c echo.Context) error {
	// Get user by ID or email
	id := c.QueryParam("id")
	email := c.QueryParam("email")

	if id == "" && email == "" {
		return response.ValidationError(c, "Either id or email must be provided")
	}

	var user *model.User
	var err error

	// if id is not empty, get user by id
	if id != "" {
		if _, err := uuid.Parse(id); err != nil {
			return response.ValidationError(c, "Invalid user ID format")
		}
		user, err = h.userService.GetUserByID(c.Request().Context(), id)
		if err != nil {
			if err == constants.ErrNotFound {
				return response.NotFound(c, "User not found")
			}
			return response.InternalError(c)
		}
	}

	// if email is not empty, get user by email
	if email != "" {
		user, err = h.userService.GetUserByEmail(c.Request().Context(), email)
		if err != nil {
			if err == constants.ErrNotFound {
				return response.NotFound(c, "User not found")
			}
			return response.InternalError(c)
		}

	}
	return response.Success(c, http.StatusOK, user)
}

// @Summary Update user
// @Description Updates an existing user's information (Admin and Instructor only)
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User UUID" format(uuid)
// @Param user body model.User true "Updated user object"
// @Success 200 {object} response.Response{data=model.User} "User updated successfully"
// @Failure 400 {object} response.Response "Invalid UUID format or request body"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden - Admin or Instructor role required"
// @Failure 404 {object} response.Response "User not found"
// @Failure 500 {object} response.Response "Internal server error"
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
// @Description Deletes a user by their UUID (Admin only)
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User UUID" format(uuid)
// @Success 200 {object} response.Response "User deleted successfully"
// @Failure 400 {object} response.Response "Invalid UUID format"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden - Admin role required"
// @Failure 404 {object} response.Response "User not found"
// @Failure 500 {object} response.Response "Internal server error"
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

// @Summary Invite users
// @Description Invites multiple users to the platform by creating their accounts and sending credentials via email
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param invite body model.UserInviteRequest true "User invite request"
// @Success 200 {object} response.Response "Users invited successfully"
// @Failure 400 {object} response.Response "Invalid request payload"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden - Not an admin"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/invite [post]
func (h *UserHandler) InviteUsers(c echo.Context) error {
	// // Get user role from context (set by auth middleware)
	// role := c.Get("userRole")
	// if role != constants.AdminRole {
	// 	return response.Error(c, http.StatusForbidden, "Only admins can invite users")
	// }

	var req model.UserInviteRequest
	if err := c.Bind(&req); err != nil {
		return response.ValidationError(c, "Invalid request payload")
	}

	if err := h.userService.InviteUsers(c.Request().Context(), req.Users); err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, "Users invited successfully")
}

// @Summary Request password reset
// @Description Sends a password reset email to the user
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body model.PasswordResetRequest true "Password reset request"
// @Success 200 {object} response.Response "Password reset email sent"
// @Failure 400 {object} response.Response "Invalid request payload"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/request-reset-password [post]
func (h *UserHandler) RequestPasswordReset(c echo.Context) error {
	var req model.PasswordResetRequest
	if err := c.Bind(&req); err != nil {
		return response.ValidationError(c, "Invalid request payload")
	}

	if err := h.userService.RequestPasswordReset(c.Request().Context(), req.Email); err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, "If the email exists, a password reset link has been sent")
}

// @Summary Reset password
// @Description Resets user's password using the reset token
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body model.PasswordResetConfirm true "Password reset confirmation"
// @Success 200 {object} response.Response "Password reset successful"
// @Failure 400 {object} response.Response "Invalid request payload or token"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/reset-password [post]
func (h *UserHandler) ResetPassword(c echo.Context) error {
	var req model.PasswordResetConfirm
	if err := c.Bind(&req); err != nil {
		return response.ValidationError(c, "Invalid request payload")
	}

	// Validate the request payload
	if err := c.Validate(req); err != nil {
		return response.ValidationError(c, "Invalid request payload")
	}

	if err := h.userService.ResetPassword(c.Request().Context(), req.Token, req.NewPassword); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid or expired reset token")
	}

	return response.Success(c, http.StatusOK, "Password reset successful")
}
