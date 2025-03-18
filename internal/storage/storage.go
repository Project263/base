package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"theAesthetics.ru/base/config"
)

func InitPostgres(cfg *config.Config) error {
	ctx := context.Background()
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Ошибка загрузки .env")
	}

	pool, err := pgxpool.New(ctx, cfg.POSTGRES_DSN)

	err = pool.Ping(ctx)

	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}

	defer pool.Close()
	InitTables(pool)

	return nil
}
