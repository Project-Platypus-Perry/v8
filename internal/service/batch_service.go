package service

import (
	"context"

	"github.com/project-platypus-perry/v8/internal/model"
	"github.com/project-platypus-perry/v8/internal/repository"
)

type BatchService interface {
	CreateBatch(ctx context.Context, batch *model.Batch, userBatch *model.UsersBatches) error
	GetBatch(ctx context.Context, id string, userID string, organizationID string) (*model.BatchResponseModel, error)
	UpdateBatch(ctx context.Context, batch *model.Batch) error
	DeleteBatch(ctx context.Context, id string) error
	AddUserToBatch(ctx context.Context, batchID string, userIDs []string, organizationID string) error
	RemoveUserFromBatch(ctx context.Context, batchID string, userIDs []string) error
	ListUserBatches(ctx context.Context, userID string, organizationID string) ([]*model.BatchResponseModel, error)
}

type batchService struct {
	batchRepo repository.BatchRepository
}

func NewBatchService(batchRepo repository.BatchRepository) BatchService {
	return &batchService{batchRepo: batchRepo}
}

func (s *batchService) CreateBatch(ctx context.Context, batch *model.Batch, userBatch *model.UsersBatches) error {
	return s.batchRepo.CreateBatch(ctx, batch, userBatch)
}

func (s *batchService) GetBatch(ctx context.Context, id string, userID string, organizationID string) (*model.BatchResponseModel, error) {
	return s.batchRepo.GetBatch(ctx, id, userID, organizationID)
}

func (s *batchService) UpdateBatch(ctx context.Context, batch *model.Batch) error {
	return s.batchRepo.UpdateBatch(ctx, batch)
}

func (s *batchService) DeleteBatch(ctx context.Context, id string) error {
	return s.batchRepo.DeleteBatch(ctx, id)
}

func (s *batchService) AddUserToBatch(ctx context.Context, batchID string, userIDs []string, organizationID string) error {
	// // if user in batch
	// userInBatch, err := s.batchRepo.IsUserInBatch(ctx, userIDs[0], batchID, false)
	// if err != nil {
	// 	return err
	// }
	// if userInBatch {
	// 	return errors.New("user already in batch")
	// }
	return s.batchRepo.AddUserToBatch(ctx, batchID, userIDs, organizationID)
}
func (s *batchService) RemoveUserFromBatch(ctx context.Context, batchID string, userIDs []string) error {
	// if user in batch
	// userInBatch, err := s.batchRepo.IsUserInBatch(ctx, userIDs[0], batchID, false)
	// if err != nil {
	// 	return err
	// }
	// if !userInBatch {
	// 	return errors.New("user not in batch")
	// }
	return s.batchRepo.RemoveUserFromBatch(ctx, batchID, userIDs)
}

func (s *batchService) ListUserBatches(ctx context.Context, userID string, organizationID string) ([]*model.BatchResponseModel, error) {
	return s.batchRepo.ListUserBatches(ctx, userID, organizationID)
}
