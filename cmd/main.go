package main

import (
	config2 "task-manager/internal/infrastructure/config"
	"task-manager/internal/infrastructure/database"
	"task-manager/internal/infrastructure/lib/logger"
	"task-manager/internal/infrastructure/server"
)

func main() {
	appConfig := config2.NewAppConfig()
	loadDB()
	logger.SetupLogger(appConfig.Env)

	server.New(appConfig)
}

func loadDB() {
	dbConfig := config2.NewDBConfig()
	database.InitDB(*dbConfig)
}
