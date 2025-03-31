package config

import (
	"github.com/caarlos0/env"
)

type Config struct {
	POSTGRES_DSN string `env:"POSTGRES_DSN" envDefault:"postgresql://root:123@localhost:5432/base?sslmode=disable"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
