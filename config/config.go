package config

import (
	"github.com/caarlos0/env"
)

type Config struct {
	Host         string `env:"HOST" envDefault:"localhost"`
	Port         string `env:"PORT" envDefault:"8080"`
	POSTGRES_DSN string `env:"POSTGRES_DSN" envDefault:"postgresql://root:123@localhost:5432/base?sslmode=disable"`
	LogLevel     string `env:"LOG_LEVEL" envDefault:"debug"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
