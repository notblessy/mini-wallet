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
	ErrAlreadyDisabled      = errors.New("already disabled")
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
	route.Use(middleware.Logger())
	route.Use(middleware.Recover())

	routes := route.Group("/api/v1")
	routes.Use(middleware.JWTWithConfig(mdw.JWTConfig()))

	routes.POST("/wallet", h.enableWalletHandler)
	routes.GET("/wallet", h.viewBalanceHandler)
	routes.PATCH("/wallet", h.disableWalletHandler)

	routes.POST("/deposits", h.depositHandler)
	routes.POST("/withdrawals", h.withdrawalHandler)

}
