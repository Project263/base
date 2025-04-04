package main

import (
	"base/config"
	"base/internal/database"
	"base/internal/logger"
	"base/internal/middlewares"
	"base/internal/router"
	"context"

	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	// init config
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	// init logger
	logger.InitLogger(cfg.LOG_LEVEL, cfg.MODE)
	// init database
	pool := database.ConnectDB(cfg, ctx)

	// init echo
	e := echo.New()

	aMiddleware := middlewares.NewAuthMiddleware(cfg)
	e.Use(aMiddleware.CheckAuthToken)

	router.InitRouter(e, pool)
	e.Logger.Fatal(e.Start(":3000"))

	// graceful shotdown
}
