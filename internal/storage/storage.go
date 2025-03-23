package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"theaesthetics.ru/base/config"
)

func InitPostgres(cfg *config.Config, ctx context.Context, log *logrus.Logger) *pgxpool.Pool {
	pool, err := pgxpool.New(ctx, cfg.POSTGRES_DSN)

	if err != nil {
		log.Errorf("ошибка подключения к базе данных: %s", err.Error())
		return nil
	}

	err = pool.Ping(ctx)

	if err != nil {
		log.Errorf("ошибка подключения к базе данных: %s", err.Error())
		return nil
	}

	return pool
}
