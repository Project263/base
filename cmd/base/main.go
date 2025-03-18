package main

import (
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
	// run http server
}
