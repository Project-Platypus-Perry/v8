package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/project-platypus-perry/v8/internal/model"
	"github.com/project-platypus-perry/v8/internal/service"
	"github.com/project-platypus-perry/v8/pkg/logger"
	"github.com/project-platypus-perry/v8/pkg/response"
	"go.uber.org/zap"
)

type BatchHandler struct {
	batchService service.BatchService
}

func NewBatchHandler(batchService service.BatchService) *BatchHandler {
	return &BatchHandler{batchService: batchService}
}

// @Summary Create a new batch
// @Description Create a new batch (Admin only)
// @Tags batch
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param batch body model.Batch true "Batch details"
// @Success 201 {object} response.Response "Batch created successfully"
// @Failure 400 {object} response.Response "Invalid request payload or validation error"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden - Admin role required"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /batch [post]
func (h *BatchHandler) CreateBatch(c echo.Context) error {
	batch := new(model.Batch)
	if err := c.Bind(batch); err != nil {
		return response.ValidationError(c, err.Error())
	}

	logger.Info("Context", zap.Any("context", c.Request().Context()))

	// get organization ID from context
	organizationID := c.Get("OrganizationID").(string)
	batch.OrganizationID = organizationID

	userBatch := &model.UsersBatches{
		UserID:         c.Get("UserID").(string),
		BatchID:        batch.ID,
		OrganizationID: organizationID,
	}

	if err := c.Validate(batch); err != nil {
		return response.ValidationError(c, err.Error())
	}

	if err := h.batchService.CreateBatch(c.Request().Context(), batch, userBatch); err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusCreated, "Batch created successfully")
}

// @Summary Get batch by ID
// @Description Retrieve a specific batch by ID
// @Tags batch
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Batch ID"
// @Success 200 {object} response.Response{data=model.BatchResponseModel} "Batch found successfully"
// @Failure 400 {object} response.Response "Invalid batch ID"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 404 {object} response.Response "Batch not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /batch/{id} [get]
func (h *BatchHandler) GetBatch(c echo.Context) error {
	id := c.Param("id")

	userID := c.Get("UserID").(string)
	organizationID := c.Get("OrganizationID").(string)

	batch, err := h.batchService.GetBatch(c.Request().Context(), id, userID, organizationID)
	if err != nil {
		return response.Error(c, http.StatusNotFound, err.Error())
	}

	return response.Success(c, http.StatusOK, batch)
}

// @Summary Add users to batch
// @Description Add multiple users to a batch (Admin only)
// @Tags batch
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body model.AssociateUserToBatchRequest true "Users to add to batch"
// @Success 200 {object} response.Response "User added to batch successfully"
// @Failure 400 {object} response.Response "Invalid request payload or validation error"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden - Admin role required"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /batch/users/add [post]
func (h *BatchHandler) AddUserToBatch(c echo.Context) error {
	var request model.AssociateUserToBatchRequest
	if err := c.Bind(&request); err != nil {
		return response.ValidationError(c, err.Error())
	}

	// validate request
	if err := c.Validate(request); err != nil {
		return response.ValidationError(c, err.Error())
	}

	organizationID := c.Get("OrganizationID").(string)

	if err := h.batchService.AddUserToBatch(c.Request().Context(), request.BatchID, request.UserIDs, organizationID); err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, "User added to batch successfully")
}

// @Summary Remove users from batch
// @Description Remove multiple users from a batch (Admin only)
// @Tags batch
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body model.AssociateUserToBatchRequest true "Users to remove from batch"
// @Success 200 {object} response.Response "User removed from batch successfully"
// @Failure 400 {object} response.Response "Invalid request payload or validation error"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden - Admin role required"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /batch/users/remove [post]
func (h *BatchHandler) RemoveUserFromBatch(c echo.Context) error {
	var request model.AssociateUserToBatchRequest

	if err := c.Bind(&request); err != nil {
		return response.ValidationError(c, err.Error())
	}

	if err := c.Validate(request); err != nil {
		return response.ValidationError(c, err.Error())
	}

	if err := h.batchService.RemoveUserFromBatch(c.Request().Context(), request.BatchID, request.UserIDs); err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, "User removed from batch successfully")
}

// @Summary List user's batches
// @Description Get all batches that the authenticated user belongs to
// @Tags batch
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=[]model.BatchResponseModel} "User batches retrieved successfully"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /batch/list [get]
func (h *BatchHandler) ListUserBatches(c echo.Context) error {
	userID := c.Get("UserID").(string)
	organizationID := c.Get("OrganizationID").(string)
	batches, err := h.batchService.ListUserBatches(c.Request().Context(), userID, organizationID)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, batches)
}
