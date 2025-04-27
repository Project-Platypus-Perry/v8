package service

import (
	"context"

	"github.com/project-platypus-perry/v8/internal/model"
	"github.com/project-platypus-perry/v8/internal/repository"
	"github.com/project-platypus-perry/v8/pkg/logger"
	"go.uber.org/zap"
)

type OrganizationService interface {
	CreateOrganization(ctx context.Context, organization *model.Organization) (string, error)
	GetOrganizationByID(ctx context.Context, id string) (*model.Organization, error)
}

type organizationService struct {
	organizationRepository repository.OrganizationRepository
}

func NewOrganizationService(organizationRepository repository.OrganizationRepository) OrganizationService {
	return &organizationService{organizationRepository: organizationRepository}
}

func (s *organizationService) CreateOrganization(ctx context.Context, organization *model.Organization) (string, error) {
	logger.Info("Creating organization", zap.Any("organization", organization))

	if err := s.organizationRepository.CreateOrganization(ctx, organization); err != nil {
		logger.Error("Failed to create organization", zap.Error(err))
		return "", err
	}

	logger.Info("Organization created successfully", zap.String("id", organization.ID))
	return organization.ID, nil
}

func (s *organizationService) GetOrganizationByID(ctx context.Context, id string) (*model.Organization, error) {
	return s.organizationRepository.GetOrganizationByID(ctx, id)
}
