package service

import (
	"context"

	"github.com/project-platypus-perry/v8/internal/config"
	"github.com/project-platypus-perry/v8/internal/constants"
	"github.com/project-platypus-perry/v8/internal/model"
	"github.com/project-platypus-perry/v8/pkg/jwt"
	"github.com/project-platypus-perry/v8/pkg/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, email string, password string) (*model.User, *jwt.TokenPair, error)
	RegisterOrganization(ctx context.Context, user *model.User, organization *model.Organization) error
}

type authService struct {
	userService         UserService
	organizationService OrganizationService
	config              *config.JWTConfig
}

func NewAuthService(userService UserService, organizationService OrganizationService, config *config.JWTConfig) AuthService {
	return &authService{
		userService:         userService,
		organizationService: organizationService,
		config:              config,
	}
}

func (s *authService) Login(ctx context.Context, email string, password string) (*model.User, *jwt.TokenPair, error) {
	user, err := s.userService.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, nil, err
	}

	logger.Info("User found", zap.Any("user", user))

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, nil, err
	}

	tokenPair, err := jwt.GenerateTokenPair(user.ID, user.Role, s.config)
	if err != nil {
		return nil, nil, err
	}

	logger.Info("Token pair generated", zap.Any("tokenPair", tokenPair))

	return user, tokenPair, nil
}

func (s *authService) RegisterOrganization(ctx context.Context, user *model.User, organization *model.Organization) error {
	organizationID, err := s.organizationService.CreateOrganization(ctx, organization)
	if err != nil {
		return err
	}

	user.OrganizationID = organizationID
	user.Role = constants.RoleAdmin

	if err := s.userService.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}
