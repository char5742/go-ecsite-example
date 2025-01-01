package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	AppName   string `env:"APP_NAME" envDefault:"ecsite-sample"`
	Port      string `env:"PORT" envDefault:"8050"`
	DebugMode bool   `env:"DEBUG_MODE" envDefault:"false"`
	Database  struct {
		Host     string `env:"DB_HOST" envDefault:"localhost"`
		Port     string `env:"DB_PORT" envDefault:"5435"`
		User     string `env:"DB_USER" envDefault:"postgres"`
		Password string `env:"DB_PASSWORD" envDefault:"postgres"`
		Name     string `env:"DB_NAME" envDefault:"app"`
	}
}

var cfg Config

func LoadConfig() *Config {
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to parse environment variables: %v", err)
	}
	return &cfg
}

func GetConfig() *Config {
	if cfg.AppName == "" {
		LoadConfig()
	}
	return &cfg
}
