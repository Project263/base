package main

import (
	"theAesthetics.ru/base/internal/logger"
	"theAesthetics.ru/base/internal/storage"
)

func main() {
	// init logger
	logger.InitLogger()
	// init db
	DB := storage.InitPostgres()
	_ = DB
	// run http server
}
