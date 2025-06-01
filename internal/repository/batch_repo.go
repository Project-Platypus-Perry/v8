package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/project-platypus-perry/v8/internal/model"
	"gorm.io/gorm"
)

type BatchRepository interface {
	CreateBatch(ctx context.Context, batch *model.Batch, userBatch *model.UsersBatches) error
	GetBatch(ctx context.Context, id string, userID string, organizationID string) (*model.BatchResponseModel, error)
	UpdateBatch(ctx context.Context, batch *model.Batch) error
	DeleteBatch(ctx context.Context, id string) error
	AddUserToBatch(ctx context.Context, batchID string, userIDs []string, organizationID string) error
	RemoveUserFromBatch(ctx context.Context, batchID string, userIDs []string) error
	IsUserInBatch(ctx context.Context, userID string, batchID string, isDeleted bool) (bool, error)
	ListUserBatches(ctx context.Context, userID string, organizationID string) ([]*model.BatchResponseModel, error)
}

type batchRepository struct {
	db *gorm.DB
}

func NewBatchRepository(db *gorm.DB) BatchRepository {
	return &batchRepository{db: db}
}

func (r *batchRepository) CreateBatch(ctx context.Context, batch *model.Batch, userBatch *model.UsersBatches) error {
	batch.ID = uuid.New().String()
	userBatch.BatchID = batch.ID

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(batch).Error; err != nil {
			return err
		}
		if err := tx.Create(userBatch).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}

func (r *batchRepository) GetBatch(ctx context.Context, id string, userID string, organizationID string) (*model.BatchResponseModel, error) {
	var batch *model.BatchResponseModel
	// check if user is in batch
	var userBatch *model.UsersBatches
	if err := r.db.Where("user_id = ? AND batch_id = ? AND organization_id = ?", userID, id, organizationID).First(&userBatch).Error; err != nil {
		return nil, err
	}
	if err := r.db.Where("id = ?", id).First(&batch).Error; err != nil {
		return nil, err
	}
	return batch, nil
}

func (r *batchRepository) UpdateBatch(ctx context.Context, batch *model.Batch) error {
	return r.db.Save(batch).Error
}

func (r *batchRepository) DeleteBatch(ctx context.Context, id string) error {
	return r.db.Delete(&model.Batch{}, id).Error
}

func (r *batchRepository) AddUserToBatch(ctx context.Context, batchID string, userIDs []string, organizationID string) error {
	for _, userID := range userIDs {
		// check if any record exists for the user and batch
		var userBatch *model.UsersBatches
		if err := r.db.Where("user_id = ? AND batch_id = ? AND deleted_at is NOT NULL", userID, batchID).First(&userBatch).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := r.db.Create(&model.UsersBatches{
					UserID:         userID,
					BatchID:        batchID,
					OrganizationID: organizationID,
				}).Error; err != nil {
					return err
				}
			} else if userBatch.DeletedAt.Valid {
				if err := r.db.Model(&userBatch).Update("deleted_at", nil).Error; err != nil {
					return err
				}
			} else {
				return errors.New("user already in batch")
			}
		}
	}
	return nil
}

func (r *batchRepository) RemoveUserFromBatch(ctx context.Context, batchID string, userIDs []string) error {
	for _, userID := range userIDs {
		if err := r.db.Delete(&model.UsersBatches{}, "user_id = ? AND batch_id = ?", userID, batchID).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *batchRepository) ListUserBatches(ctx context.Context, userID string, organizationID string) ([]*model.BatchResponseModel, error) {
	var batches []*model.BatchResponseModel
	if err := r.db.Joins("JOIN users_batches ON batches.id = users_batches.batch_id").
		Where("users_batches.user_id = ? AND users_batches.organization_id = ? AND users_batches.deleted_at IS NULL", userID, organizationID).
		Find(&batches).Error; err != nil {
		return nil, err
	}
	return batches, nil
}

func (r *batchRepository) IsUserInBatch(ctx context.Context, userID string, batchID string, isDeleted bool) (bool, error) {
	var userBatch *model.UsersBatches
	if err := r.db.Where("user_id = ? AND batch_id = ? AND deleted_at = ?", userID, batchID, isDeleted).First(&userBatch).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
