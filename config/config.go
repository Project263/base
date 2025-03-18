package config

import "github.com/caarlos0/env"

type Config struct {
	Host         string `env:"HOST" envDefault:"9090"`
	Port         string `env:"PORT" envDefault:"debug"` //info
	POSTGRES_DSN string `env:"POSTGRES_DSN" envDefault:"postgresql://user3:password1@localhost:5433/pool1?sslmode=disable"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
