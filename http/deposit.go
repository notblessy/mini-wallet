package http

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/notblessy/mini-wallet/model"
	"github.com/notblessy/mini-wallet/utils"
	log "github.com/sirupsen/logrus"
)

// depositHandler :nodoc:
func (h *HTTPService) depositHandler(c echo.Context) error {
	var data model.Deposit

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, utils.DefaultResponse{
			Status: utils.RespStatusFail,
			Data:   err.Error(),
		})
	}
	if err := c.Validate(&data); err != nil {
		return c.JSON(http.StatusBadRequest, utils.DefaultResponse{
			Status: utils.RespStatusFail,
			Data:   err.Error(),
		})
	}

	logger := log.WithFields(log.Fields{
		"context": utils.Encode(c),
		"request": utils.Encode(data),
	})

	deposit, err := h.depositRepo.FindByReference(&data.ReferenceID)
	if deposit != nil {
		logger.Error(ErrDuplicateReferenceID)
		return c.JSON(http.StatusInternalServerError, utils.ResponseError{
			Status:  utils.RespStatusFail,
			Message: ErrDuplicateReferenceID.Error(),
		})
	}

	now := time.Now()
	data.ID = uuid.New().String()
	data.DepositedAt = &now
	data.Status = model.DepositStatus_Success
	data.DepositedBy = "ea0212d3-abd6-406f-8c67-23d8e814a2436"

	wallet, err := h.walletRepo.FindByOwner(&data.DepositedBy)
	if err != nil {
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, utils.ResponseError{
			Status:  utils.RespStatusError,
			Message: err.Error(),
		})
	}

	wallet.Balance = wallet.Balance + data.Amount

	err = h.depositRepo.Create(&data, wallet)
	if err != nil {
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, utils.ResponseError{
			Status:  utils.RespStatusError,
			Message: err.Error(),
		})
	}

	depositResp := data.NewResponse()

	return c.JSON(http.StatusOK, utils.DefaultResponse{
		Status: utils.RespStatusSuccess,
		Data:   depositResp,
	})
}
