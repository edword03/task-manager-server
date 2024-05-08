package app

import (
	"task-manager/internal/config"
	"task-manager/internal/database/postgres"
	"task-manager/internal/database/redis"
	"task-manager/internal/server"
)

func Run() {
	appConfig := config.NewAppConfig()
	loadDB()

	server.New(appConfig)
}

func loadDB() {
	dbConfig := config.NewDBConfig()
	postgres.InitDB(dbConfig)

	redis.InitRedis(dbConfig)
}
