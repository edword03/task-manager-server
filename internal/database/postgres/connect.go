package postgres

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"task-manager/internal/config"
	"task-manager/internal/database/postgres/models"
)

var Db *gorm.DB

func InitDB(cfg *config.DBConfig) *gorm.DB {
	Db = connectDB(cfg)
	return Db
}

func connectDB(cfg *config.DBConfig) *gorm.DB {
	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", "localhost", cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Connection error", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Workspace{}, &models.Tag{}, &models.Task{})
	if err != nil {
		log.Fatal("Migration error", err)
	}

	return db
}
