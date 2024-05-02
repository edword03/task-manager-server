package main

import (
	"task-manager/internal/infrastructure/config"
	"task-manager/internal/infrastructure/database/postgres"
	"task-manager/internal/infrastructure/database/redis"
	"task-manager/internal/infrastructure/server"
)

func main() {
	appConfig := config.NewAppConfig()
	loadDB()

	server.New(appConfig)
}

func loadDB() {
	dbConfig := config.NewDBConfig()
	postgres.InitDB(dbConfig)

	redis.InitRedis(dbConfig)
}
