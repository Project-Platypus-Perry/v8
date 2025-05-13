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

func (h *BatchHandler) ListUserBatches(c echo.Context) error {
	userID := c.Get("UserID").(string)
	organizationID := c.Get("OrganizationID").(string)
	batches, err := h.batchService.ListUserBatches(c.Request().Context(), userID, organizationID)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, batches)
}
