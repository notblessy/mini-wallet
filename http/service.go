package http

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	mdw "github.com/notblessy/mini-wallet/middleware"
	"github.com/notblessy/mini-wallet/model"

	"gorm.io/gorm"
)

var (
	// ErrBadRequest :nodoc:
	ErrBadRequest = errors.New("bad request")

	// ErrIncorrectEmailOrPassword :nodoc:
	ErrIncorrectEmailOrPassword = errors.New("incorrect email or password")

	// ErrNotFound :nodoc:
	ErrNotFound = errors.New("not found")
)

// HTTPService :nodoc:
type HTTPService struct {
	userRepo model.UserRepository
	db       gorm.DB
}

// NewHTTPService :nodoc:
func NewHTTPService() *HTTPService {
	return new(HTTPService)
}

// RegisterUserRepository :nodoc:
func (h *HTTPService) RegisterUserRepository(u model.UserRepository) {
	h.userRepo = u
}

// Routes :nodoc:
func (h *HTTPService) Routes(route *echo.Echo) {
	route.POST("api/v1/init", h.initHandler)

	routes := route.Group("/api/v1")
	routes.Use(middleware.Logger())
	routes.Use(middleware.Recover())
	routes.Use(middleware.JWTWithConfig(mdw.JWTConfig()))
}
