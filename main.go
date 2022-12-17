package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/notblessy/mini-wallet/config"
	"github.com/notblessy/mini-wallet/db"
)

func main() {
	initDB := db.InitiateMysql()
	defer db.CloseMysql(initDB)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	log.Fatal(e.Start(":" + config.HTTPPort()))
}
