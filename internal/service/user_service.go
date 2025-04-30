package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/project-platypus-perry/v8/internal/constants"
	"github.com/project-platypus-perry/v8/internal/model"
	"github.com/project-platypus-perry/v8/internal/repository"
	emailService "github.com/project-platypus-perry/v8/pkg/email_service"
	"github.com/project-platypus-perry/v8/pkg/logger"
	"github.com/project-platypus-perry/v8/pkg/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateUser(ctx context.Context, id string, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, id string) error
	InviteUsers(ctx context.Context, invites []model.UserInvite) error
	RequestPasswordReset(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, token string, newPassword string) error
}

type userService struct {
	repo         repository.UserRepository
	emailService emailService.EmailService
	jwtSecret    string
}

func NewUserService(repo repository.UserRepository, emailService emailService.EmailService, jwtSecret string) UserService {
	return &userService{
		repo:         repo,
		emailService: emailService,
		jwtSecret:    jwtSecret,
	}
}

func (s *userService) CreateUser(ctx context.Context, user *model.User) error {
	// Encrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repo.CreateUser(ctx, user)
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}

func (s *userService) UpdateUser(ctx context.Context, id string, user *model.User) (*model.User, error) {
	return s.repo.UpdateUser(ctx, id, user)
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *userService) InviteUsers(ctx context.Context, invites []model.UserInvite) error {
	for _, invite := range invites {
		// Check if user already exists
		if _, err := s.GetUserByEmail(ctx, invite.Email); err == nil {
			return errors.New("user with email " + invite.Email + " already exists")
		}

		// Generate password
		password, err := utils.GeneratePassword(12)
		if err != nil {
			return err
		}

		// Create user
		user := &model.User{
			ID:             uuid.New().String(),
			OrganizationID: invite.OrganizationID,
			Role:           invite.Role,
			Name:           invite.Name,
			Email:          invite.Email,
			Password:       password,
			Phone:          invite.Phone,
		}

		if err := s.CreateUser(ctx, user); err != nil {
			return err
		}

		// Send invite email
		if err := s.emailService.SendInviteEmail(invite.Email, invite.Name, invite.Email, password); err != nil {
			// If email fails, we should probably log this and handle it appropriately
			// For now, we'll continue with other invites
			logger.Error("Failed to send invite email", zap.Error(err))
			continue
		}
	}

	return nil
}

func (s *userService) RequestPasswordReset(ctx context.Context, email string) error {
	user, err := s.GetUserByEmail(ctx, email)
	if err != nil {
		if err == constants.ErrNotFound {
			// Don't reveal if email exists or not
			return nil
		}
		return err
	}

	// Generate reset token
	token, err := utils.GeneratePasswordResetToken(user.ID, s.jwtSecret, 24) // 24 hours expiry
	if err != nil {
		return err
	}

	// Send reset email
	return s.emailService.SendPasswordResetEmail(email, token)
}

func (s *userService) ResetPassword(ctx context.Context, token string, newPassword string) error {
	// Validate token
	userID, err := utils.ValidatePasswordResetToken(token, s.jwtSecret)
	if err != nil {
		return err
	}

	// Get user
	user, err := s.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update password
	user.Password = string(hashedPassword)
	_, err = s.UpdateUser(ctx, userID, user)
	return err
}
