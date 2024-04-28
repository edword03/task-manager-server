package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"task-manager/internal/infrastructure/config"
)

func New(cfg *config.AppConfig) {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(cors.Default())
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	log.Info("Server starting...")

	if err := r.Run(cfg.HTTPServer.Address); err != nil {
		log.Fatal(err)
	}
}
