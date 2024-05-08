package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"task-manager/internal/config"
	"task-manager/internal/controllers/http/auth"
	"task-manager/internal/database/postgres"
	userRepo "task-manager/internal/database/postgres/user"
	"task-manager/internal/database/redis"
	redisRepo "task-manager/internal/database/redis/repositories"
	"task-manager/internal/domain/services"
	"task-manager/internal/pkg/logger"
)

func New(cfg *config.AppConfig) {
	r := gin.Default()

	logger.SetupLogger(cfg.Env)

	r.Use(gin.Logger())
	r.Use(cors.Default())
	r.Use(gin.Recovery())

	userRepository := userRepo.NewUserRepo(postgres.Db)
	authService := services.NewAuthService(userRepository)

	tokenRepo := redisRepo.NewTokenRepo(redis.TokensClient)
	tokenService := auth.NewJWTService(cfg, tokenRepo)

	auth.NewAuthController(r, authService, tokenService, cfg)

	log.Info("Server starting...")

	if err := r.Run(cfg.HTTPServer.Address); err != nil {
		log.Fatal(err)
	}
}
