package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"task-manager/internal/config"
	"task-manager/internal/controllers/http/auth"
	"task-manager/internal/controllers/http/jwt"
	"task-manager/internal/controllers/http/users"
	"task-manager/internal/database/postgres"
	userRepo "task-manager/internal/database/postgres/user"
	"task-manager/internal/database/redis"
	redisRepository "task-manager/internal/database/redis/repositories"
	"task-manager/internal/domain/services"
	"task-manager/internal/pkg/logger"
	"time"
)

func New(cfg *config.AppConfig) {
	r := gin.Default()

	logger.SetupLogger(cfg.Env)

	r.Use(gin.Logger())
	r.Use(cors.Default())
	r.Use(gin.Recovery())

	userRepository := userRepo.NewUserRepo(postgres.Db)
	authService := services.NewAuthService(userRepository)
	userService := services.NewUserService(userRepository)

	tokenRepo := redisRepository.NewRedisRepo(redis.TokensClient)
	tokenService := jwt.NewJWTService(cfg, tokenRepo, auth.MaxAgeCookie*time.Duration(cfg.RefreshMaxAge))

	auth.NewAuthController(r, authService, tokenService, cfg)
	users.NewUsersController(r, userService, tokenService)

	log.Info("Server starting...")

	if err := r.Run(cfg.HTTPServer.Address); err != nil {
		log.Fatal(err)
	}
}
