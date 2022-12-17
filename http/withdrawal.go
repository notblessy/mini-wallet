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

// withdrawalHandler :nodoc:
func (h *HTTPService) withdrawalHandler(c echo.Context) error {
	var data model.Withdrawal

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

	withdrawal, err := h.withdrawalRepo.FindByReference(&data.ReferenceID)
	if withdrawal != nil {
		logger.Error(ErrDuplicateReferenceID)
		return c.JSON(http.StatusInternalServerError, utils.DefaultResponse{
			Status: utils.RespStatusFail,
			Data:   ErrDuplicateReferenceID.Error(),
		})
	}

	now := time.Now()
	data.ID = uuid.New().String()
	data.WithdrawnAt = &now
	data.Status = model.WithdrawalStatus_Success
	data.WithdrawnBy = "ea0212d3-abd6-406f-8c67-23d8e814a2436"

	wallet, err := h.walletRepo.FindByOwner(&data.WithdrawnBy)
	if err != nil {
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, utils.ResponseError{
			Status:  utils.RespStatusError,
			Message: err.Error(),
		})
	}

	if wallet.Balance < data.Amount {
		logger.Error(ErrNotEnoughBalance)
		return c.JSON(http.StatusInternalServerError, utils.DefaultResponse{
			Status: utils.RespStatusFail,
			Data:   ErrNotEnoughBalance.Error(),
		})
	}

	wallet.Balance = wallet.Balance - data.Amount

	err = h.withdrawalRepo.Create(&data, wallet)
	if err != nil {
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, utils.ResponseError{
			Status:  utils.RespStatusError,
			Message: err.Error(),
		})
	}

	withdrawalResp := data.NewResponse()

	return c.JSON(http.StatusOK, utils.DefaultResponse{
		Status: utils.RespStatusSuccess,
		Data:   withdrawalResp,
	})
}
