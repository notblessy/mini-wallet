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
	ErrExisted              = errors.New("data existed")
	ErrNotEnoughBalance     = errors.New("not enough balance")
	ErrDuplicateReferenceID = errors.New("duplicate reference id")
)

// HTTPService :nodoc:
type HTTPService struct {
	userRepo       model.UserRepository
	walletRepo     model.WalletRepository
	depositRepo    model.DepositRepository
	withdrawalRepo model.WithdrawalRepository
	db             gorm.DB
}

// NewHTTPService :nodoc:
func NewHTTPService() *HTTPService {
	return new(HTTPService)
}

// RegisterUserRepository :nodoc:
func (h *HTTPService) RegisterUserRepository(u model.UserRepository) {
	h.userRepo = u
}

// RegisterWalletRepository :nodoc:
func (h *HTTPService) RegisterWalletRepository(w model.WalletRepository) {
	h.walletRepo = w
}

// RegisterDepositRepository :nodoc:
func (h *HTTPService) RegisterDepositRepository(d model.DepositRepository) {
	h.depositRepo = d
}

// RegisterWithdrawalRepository :nodoc:
func (h *HTTPService) RegisterWithdrawalRepository(w model.WithdrawalRepository) {
	h.withdrawalRepo = w
}

// Routes :nodoc:
func (h *HTTPService) Routes(route *echo.Echo) {
	route.POST("api/v1/init", h.initHandler)

	routes := route.Group("/api/v1")
	routes.Use(middleware.Logger())
	routes.Use(middleware.Recover())
	routes.Use(middleware.JWTWithConfig(mdw.JWTConfig()))

	route.POST("api/v1/wallet", h.enableWalletHandler)
	route.GET("api/v1/wallet", h.viewBalanceHandler)

	route.POST("api/v1/deposits", h.depositHandler)
	route.POST("api/v1/withdrawals", h.withdrawalHandler)

}
