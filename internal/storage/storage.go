package storage

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitPostgres() *sql.DB {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Ошибка загрузки .env")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("DBNAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err.Error())
		panic("База данных нихуя не работает")
	}

	InitTables(DB)

	return DB
}
