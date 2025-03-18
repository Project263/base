package main

import "theAesthetics.ru/base/internal/storage"

func main() {
	// init logger
	// init db
	DB := storage.InitPostgres()
	_ = DB
	// run http server
}
