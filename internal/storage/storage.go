package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"theAesthetics.ru/base/config"
)

func InitPostgres(cfg *config.Config) *pgxpool.Pool {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, cfg.POSTGRES_DSN)

	if err != nil {
		log.Fatalf("ошибка подключения к базе данных: %s", err.Error())
		return nil
	}

	err = pool.Ping(ctx)

	if err != nil {
		log.Fatalf("ошибка подключения к базе данных: %s", err.Error())
		return nil
	}

	return pool
}
