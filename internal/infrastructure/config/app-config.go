package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"time"
)

type AppConfig struct {
	Env           string        `yaml:"app_env" envDefault:"development"`
	DBPath        string        `yaml:"db_path" envDefault:"./db"`
	SecretKey     string        `yaml:"secret_key" env-required:"true"`
	AccessMaxAge  time.Duration `yaml:"accessMaxAge" envDefault:"15m"`
	RefreshMaxAge int           `yaml:"refreshMaxAge" envDefault:"10"`
	HTTPServer    `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" envDefault:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User        string        `yaml:"user" env-required:"true"`
	Password    string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

func NewAppConfig() *AppConfig {
	var configPath = "configs/config.yaml"

	var cfg AppConfig

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
