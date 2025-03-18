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
	DB := storage.InitPostgres(cfg)
	_ = DB
	// run http server
}
