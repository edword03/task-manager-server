package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type DBConfig struct {
	DBPort     string `env:"DB_PORT" envDefault:"5432"`
	DBPassword string `env:"POSTGRES_PASSWORD" env-required:"true"`
	DBName     string `env:"POSTGRES_DB"`
	DBUser     string `env:"POSTGRES_USER" envDefault:"user-service"`

	RedisHost string `env:"REDIS_HOST" envDefault:"localhost"`
	RedisPort string `env:"REDIS_PORT" envDefault:"6379"`
	RedisPass string `env:"REDIS_PASSWORD" env-required:"true"`
}

func NewDBConfig() *DBConfig {
	var cfg DBConfig

	if err := cleanenv.ReadConfig("deployments/.env", &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
