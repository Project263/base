package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	"theaesthetics.ru/base/config"
	"theaesthetics.ru/base/internal/logger"
	"theaesthetics.ru/base/internal/router"
	"theaesthetics.ru/base/internal/storage"
)

func main() {
	ctx := context.Background()
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err.Error())
	}

	log := logger.InitLogger(cfg.LogLevel)
	pool := storage.InitPostgres(cfg, ctx, log)
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	router.InitRouter(e, pool)

	go func() {
		log.Error(e.Start(":" + cfg.Port))
	}()

	log.Info("start service")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Info("stop service")

	if err := e.Shutdown(ctx); err != nil {
		logrus.Error("errors on stoping server", err)
	}

	pool.Close()
}
