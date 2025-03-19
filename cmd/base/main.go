package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	"theaesthetics.ru/base/config"
	"theaesthetics.ru/base/internal/logger"
	"theaesthetics.ru/base/internal/storage"
)

func main() {
	// init config
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err.Error())
	}

	// init logger
	logger.InitLogger(cfg.LogLevel)
	logrus.Info("test log")

	// init pool
	pool := storage.InitPostgres(cfg)
	_ = pool

	// run echo server
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	api := e.Group("/api/v1")
	api.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
