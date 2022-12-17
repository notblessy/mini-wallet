package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/notblessy/mini-wallet/config"
	"github.com/notblessy/mini-wallet/db"
	"github.com/notblessy/mini-wallet/http"
	"github.com/notblessy/mini-wallet/repository"
	"github.com/notblessy/mini-wallet/utils"
)

func main() {
	initDB := db.InitiateMysql()
	defer db.CloseMysql(initDB)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	userRepo := repository.NewUserRepository(initDB)
	walletRepo := repository.NewWalletRepository(initDB)
	depositRepo := repository.NewDepositRepository(initDB, walletRepo)
	withdrawalRepo := repository.NewWithdrawalRepository(initDB, walletRepo)

	httpSvc := http.NewHTTPService()
	httpSvc.RegisterUserRepository(userRepo)
	httpSvc.RegisterWalletRepository(walletRepo)
	httpSvc.RegisterDepositRepository(depositRepo)
	httpSvc.RegisterWithdrawalRepository(withdrawalRepo)

	httpSvc.Routes(e)

	log.Fatal(e.Start(":" + config.HTTPPort()))
}
