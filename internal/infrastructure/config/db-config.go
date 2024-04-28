package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type DBConfig struct {
	DBPort     string `env:"DB_PORT" envDefault:"5432"`
	DBPassword string `env:"POSTGRES_PASSWORD"`
	DBName     string `env:"POSTGRES_DB"`
	DBUser     string `env:"POSTGRES_USER" envDefault:"user-service"`
}

func NewDBConfig() *DBConfig {
	var cfg DBConfig

	if err := cleanenv.ReadConfig("deployments/.env", &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
