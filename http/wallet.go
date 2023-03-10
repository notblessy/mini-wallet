package http

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/notblessy/mini-wallet/middleware"
	"github.com/notblessy/mini-wallet/model"
	"github.com/notblessy/mini-wallet/utils"
	log "github.com/sirupsen/logrus"
)

// enableWalletHandler :nodoc:
func (h *HTTPService) enableWalletHandler(c echo.Context) error {
	logger := log.WithFields(log.Fields{
		"context": utils.Encode(c),
	})

	user, err := middleware.GetSessionClaims(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ResponseError{
			Status:  utils.RespStatusError,
			Message: err.Error(),
		})
	}

	now := time.Now()

	data := model.Wallet{
		ID:         uuid.New().String(),
		OwnedBy:    user.CustomerXid,
		Status:     model.WalletStatus_Enabled,
		Balance:    0,
		EnabledAt:  &now,
		DisabledAt: nil,
	}

	wallet, err := h.walletRepo.FindByOwner(&data.OwnedBy)
	if err != nil {
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, utils.ResponseError{
			Status:  utils.RespStatusError,
			Message: err.Error(),
		})
	}

	if wallet != nil && wallet.Status == model.WalletStatus_Enabled {
		logger.Error(ErrExisted)
		return c.JSON(http.StatusInternalServerError, utils.DefaultResponse{
			Status: utils.RespStatusFail,
			Data:   ErrExisted.Error(),
		})
	}

	switch wallet.Status {
	case model.WalletStatus_Disabled:
		data.ID = wallet.ID
		data.OwnedBy = wallet.OwnedBy
		data.Status = model.WalletStatus_Enabled
		data.Balance = wallet.Balance
		data.EnabledAt = &now
		data.DisabledAt = nil

		err = h.walletRepo.ToggleStatus(wallet)
		if err != nil {
			logger.Error(err)
			return c.JSON(http.StatusInternalServerError, utils.DefaultResponse{
				Status: utils.RespStatusFail,
				Data:   err.Error(),
			})
		}
	default:
		err = h.walletRepo.Create(&data)
		if err != nil {
			logger.Error(err)
			return c.JSON(http.StatusInternalServerError, utils.ResponseError{
				Status:  utils.RespStatusError,
				Message: err.Error(),
			})
		}
	}

	walletResp := data.NewResponse()

	return c.JSON(http.StatusOK, utils.DefaultResponse{
		Status: utils.RespStatusSuccess,
		Data:   walletResp,
	})
}

// viewBalanceHandler :nodoc:
func (h *HTTPService) viewBalanceHandler(c echo.Context) error {
	logger := log.WithFields(log.Fields{
		"context": utils.Encode(c),
	})

	user, err := middleware.GetSessionClaims(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ResponseError{
			Status:  utils.RespStatusError,
			Message: err.Error(),
		})
	}

	data, err := h.walletRepo.FindByOwner(&user.CustomerXid)
	if err != nil {
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, utils.ResponseError{
			Status:  utils.RespStatusError,
			Message: err.Error(),
		})
	}

	walletResp := data.NewResponse()

	return c.JSON(http.StatusOK, utils.DefaultResponse{
		Status: utils.RespStatusSuccess,
		Data:   walletResp,
	})
}

// disableWalletHandler :nodoc:
func (h *HTTPService) disableWalletHandler(c echo.Context) error {
	var data model.WalletStatusRequest

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
		"data":    utils.Encode(data),
	})

	// user, err := middleware.GetSessionClaims(c)
	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, utils.ResponseError{
	// 		Status:  utils.RespStatusError,
	// 		Message: err.Error(),
	// 	})
	// }

	w := model.Wallet{
		OwnedBy: "ea0212d3-abd6-406f-8c67-23d8e814a2436",
	}

	wallet, err := h.walletRepo.FindByOwner(&w.OwnedBy)
	if err != nil {
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, utils.ResponseError{
			Status:  utils.RespStatusError,
			Message: err.Error(),
		})
	}

	if wallet.Status == model.WalletStatus_Disabled {
		logger.Error(ErrAlreadyDisabled)
		return c.JSON(http.StatusInternalServerError, utils.DefaultResponse{
			Status: utils.RespStatusError,
			Data:   ErrAlreadyDisabled.Error(),
		})
	}

	switch data.IsDisabled {
	case true:
		now := time.Now()

		wallet.Status = model.WalletStatus_Disabled
		wallet.DisabledAt = &now
		wallet.EnabledAt = nil

		err = h.walletRepo.ToggleStatus(wallet)
		if err != nil {
			logger.Error(err)
			return c.JSON(http.StatusInternalServerError, utils.DefaultResponse{
				Status: utils.RespStatusFail,
				Data:   err.Error(),
			})
		}
	}

	walletResp := wallet.NewResponse()

	return c.JSON(http.StatusOK, utils.DefaultResponse{
		Status: utils.RespStatusSuccess,
		Data:   walletResp,
	})
}
