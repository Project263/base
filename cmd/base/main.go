package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"theaesthetics.ru/base/config"
	"theaesthetics.ru/base/internal/logger"
	"theaesthetics.ru/base/internal/router"
	"theaesthetics.ru/base/internal/storage"
)

func main() {
	ctx := context.Background()
	// init config
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

	// init router
	router.InitRouter(e, pool)

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
