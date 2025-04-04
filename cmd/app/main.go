package main

import (
	"base/config"
	"base/internal/database"
	"base/internal/logger"
	"context"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	// init config
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	// init logger
	logger.InitLogger(cfg.LOG_LEVEL)
	// init database
	database.ConnectDB(cfg, ctx)

	logrus.Info("Сервис auth запущен")
	logrus.Warn("Проблема с подключением к БД")
	logrus.Error("Критическая ошибка")

	// init echo
	e := echo.New()

	e.Logger.Fatal(e.Start(":3000"))

	// graceful shotdown
}
