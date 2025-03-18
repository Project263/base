package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"theAesthetics.ru/base/config"
)

func InitPostgres(cfg *config.Config) *pgxpool.Pool {
	ctx := context.Background()
	err := godotenv.Load()

	if err != nil {
		log.Println("Ошибка загрузки .env")
	}

	pool, err := pgxpool.New(ctx, cfg.POSTGRES_DSN)

	if err != nil {
		log.Fatalf("ошибка подключения к базе данных: %w", err)
		return nil
	}

	err = pool.Ping(ctx)

	if err != nil {
		log.Fatalf("ошибка подключения к базе данных: %w", err)
		return nil
	}

	defer pool.Close()
	InitTables(pool)
	return pool
}
