package main

import (
	"github.com/labstack/echo/v4"
	"theAesthetics.ru/base/config"
	"theAesthetics.ru/base/internal/storage"
)

func main() {
	// init config
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err.Error())
	}
	// init logger
	// init pool
	pool := storage.InitPostgres(cfg)
	_ = pool
	// run echo server
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
