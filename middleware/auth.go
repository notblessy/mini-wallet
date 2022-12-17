package middleware

import (
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/notblessy/mini-wallet/config"
	logger "github.com/sirupsen/logrus"
)

type Err error

var (
	ErrUnauthorized Err = errors.New("unauthorized")
)

type JWTClaims struct {
	jwt.StandardClaims
	CustomerXid string
}

func JWTConfig() middleware.JWTConfig {
	c := middleware.JWTConfig{
		Claims:     &JWTClaims{},
		SigningKey: []byte(config.JWTSecret()),
	}

	return c
}

func GetSessionClaims(c echo.Context) (*JWTClaims, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTClaims)

	if claims == nil {
		logger.Error(ErrUnauthorized)
		return nil, ErrUnauthorized
	}

	return claims, nil
}
