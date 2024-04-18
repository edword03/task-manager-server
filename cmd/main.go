package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"task-manager/internal/config"
	"task-manager/internal/database"
	"task-manager/internal/lib/logger"
)

func main() {
	loadDB()
	appConfig := config.NewAppConfig()
	logger.SetupLogger(appConfig.Env)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	log.Info("Server starting...")

	if err := r.Run(appConfig.HTTPServer.Address); err != nil {
		log.Fatal(err)
	}
}

func loadDB() {
	dbConfig := config.NewDBConfig()
	database.InitDB(*dbConfig)
}
