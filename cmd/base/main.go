package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"theaesthetics.ru/base/config"
	"theaesthetics.ru/base/internal/handlers"
	"theaesthetics.ru/base/internal/logger"
	"theaesthetics.ru/base/internal/repository"
	"theaesthetics.ru/base/internal/services"
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

	// init pool
	pool := storage.InitPostgres(cfg)
	_ = pool

	// run echo server
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	api := e.Group("/api/v1")

	eqRepo := repository.NewEquipmentRepository(pool)
	eqService := services.NewEquipmentService(eqRepo)
	eqHandler := handlers.NewEquipmentHandler(eqService)

	api.GET("/equipment", eqHandler.GetAllEquipments)

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
