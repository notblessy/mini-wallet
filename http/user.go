package http

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/notblessy/mini-wallet/config"
	"github.com/notblessy/mini-wallet/middleware"
	"github.com/notblessy/mini-wallet/model"
	"github.com/notblessy/mini-wallet/utils"
	log "github.com/sirupsen/logrus"
)

// initHandler :nodoc:
func (h *HTTPService) initHandler(c echo.Context) error {
	var data model.User

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

	data.CreatedAt = time.Now()
	err := h.userRepo.Create(&data)
	if err != nil {
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, utils.ResponseError{
			Status:  utils.RespStatusError,
			Message: err.Error(),
		})
	}

	claims := &middleware.JWTClaims{
		CustomerXid: data.CustomerXid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.JWTSecret()))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.DefaultResponse{
			Status: utils.RespStatusFail,
			Data:   err.Error(),
		})
	}

	userResp := model.UserResponse{
		Token: t,
	}

	return c.JSON(http.StatusOK, utils.DefaultResponse{
		Status: utils.RespStatusSuccess,
		Data:   userResp,
	})
}
