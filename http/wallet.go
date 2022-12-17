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

// enableWalletHandler :nodoc:
func (h *HTTPService) enableWalletHandler(c echo.Context) error {
	logger := log.WithFields(log.Fields{
		"context": utils.Encode(c),
	})

	// user, err := middleware.GetSessionClaims(c)
	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, utils.ResponseError{
	// 		Status:  utils.RespStatusError,
	// 		Message: err.Error(),
	// 	})
	// }

	now := time.Now()

	data := model.Wallet{
		ID:         uuid.New().String(),
		OwnedBy:    "ea0212d3-abd6-406f-8c67-23d8e814a2436",
		Status:     model.WalletStatus_Enabled,
		Balance:    0,
		EnabledAt:  &now,
		DisabledAt: nil,
	}

	err := h.walletRepo.Create(&data)
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

// viewBalanceHandler :nodoc:
func (h *HTTPService) viewBalanceHandler(c echo.Context) error {
	logger := log.WithFields(log.Fields{
		"context": utils.Encode(c),
	})

	// user, err := middleware.GetSessionClaims(c)
	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, utils.ResponseError{
	// 		Status:  utils.RespStatusError,
	// 		Message: err.Error(),
	// 	})
	// }

	userID := "ea0212d3-abd6-406f-8c67-23d8e814a2436"

	data, err := h.walletRepo.FindByID(&userID)
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
