package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/notblessy/mini-wallet/config"
	logrus "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitiateMysql :nodoc:
func InitiateMysql() *gorm.DB {
	err := config.LoadENV()
	if err != nil {
		logrus.Fatal(err)
	}

	logLevel := logger.Info

	if config.ENV() == "PRODUCTION" {
		logLevel = logger.Error
	}

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", config.MysqlUser(), config.MysqlPassword(), config.MysqlHost(), config.MysqlDBName())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		logrus.Fatal(fmt.Sprintf("failed to connect: %s", err))
	}

	return db
}

// CloseMysql :nodoc:
func CloseMysql(db *gorm.DB) {
	sql, err := db.DB()
	if err != nil {
		logrus.Fatal(fmt.Sprintf("failed to disconnect: %s", err))
	}

	err = sql.Close()
	if err != nil {
		logrus.Fatal(fmt.Sprintf("failed to close: %s", err))
	}
}
