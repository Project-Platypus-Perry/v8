package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/project-platypus-perry/v8/internal/model"
	"github.com/project-platypus-perry/v8/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	CreateOrganization(ctx context.Context, organization *model.Organization) error
	GetOrganizationByID(ctx context.Context, id string) (*model.Organization, error)
}

type organizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepository{db: db}
}

func (r *organizationRepository) CreateOrganization(ctx context.Context, organization *model.Organization) error {
	organization.ID = uuid.New().String()
	logger.Info("Creating organization", zap.Any("organization", organization))
	return r.db.Create(organization).Error
}

func (r *organizationRepository) GetOrganizationByID(ctx context.Context, id string) (*model.Organization, error) {
	var organization model.Organization
	if err := r.db.Where("id = ?", id).First(&organization).Error; err != nil {
		return nil, err
	}
	return &organization, nil
}
