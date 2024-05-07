package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"task-manager/internal/domain/services/auth_service"
	"task-manager/internal/infrastructure/config"
	restapi "task-manager/internal/infrastructure/controllers/rest-api/auth"
	"task-manager/internal/infrastructure/database/postgres"
	"task-manager/internal/infrastructure/database/postgres/repositories"
	"task-manager/internal/infrastructure/database/redis"
	redisRepo "task-manager/internal/infrastructure/database/redis/repositories"
	"task-manager/internal/infrastructure/lib/logger"
)

func New(cfg *config.AppConfig) {
	r := gin.Default()

	logger.SetupLogger(cfg.Env)

	r.Use(gin.Logger())
	r.Use(cors.Default())
	r.Use(gin.Recovery())

	userRepository := repositories.NewUserRepo(postgres.Db)
	authService := auth_service.NewAuthService(userRepository)

	tokenRepo := redisRepo.NewTokenRepo(redis.TokensClient)
	tokenService := restapi.NewJWTService(cfg, tokenRepo)

	restapi.NewAuthController(r, authService, tokenService, cfg)

	log.Info("Server starting...")

	if err := r.Run(cfg.HTTPServer.Address); err != nil {
		log.Fatal(err)
	}
}
