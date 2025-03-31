package main

import (
	"base/config"
	"base/internal/database"
	"base/internal/logger"
	"context"
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
	// init echo

	// graceful shotdown
}
