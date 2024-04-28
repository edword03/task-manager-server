package database

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"task-manager/internal/infrastructure/config"
)

var Db *gorm.DB

func InitDB(cfg config.DBConfig) *gorm.DB {
	Db = connectDB(cfg)
	return Db
}

func connectDB(cfg config.DBConfig) *gorm.DB {
	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", "localhost", cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal("Connection error", err)
	}

	return db
}
