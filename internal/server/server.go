package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"task-manager/internal/config"
	"task-manager/internal/controllers/http/auth"
	"task-manager/internal/controllers/http/jwt"
	"task-manager/internal/controllers/http/tags"
	"task-manager/internal/controllers/http/users"
	"task-manager/internal/database/postgres"
	"task-manager/internal/database/postgres/tag"
	userRepo "task-manager/internal/database/postgres/user"
	"task-manager/internal/database/redis"
	redisRepository "task-manager/internal/database/redis/repositories"
	"task-manager/internal/domain/services"
	"task-manager/internal/pkg/logger"
	"time"
)

func New(cfg *config.AppConfig) {
	g := gin.Default()

	logger.SetupLogger(cfg.Env)

	g.Use(gin.Logger())
	g.Use(cors.Default())
	g.Use(gin.Recovery())

	r := g.Group("/api/v1")

	userRepository := userRepo.NewUserRepo(postgres.Db)
	authService := services.NewAuthService(userRepository)
	userService := services.NewUserService(userRepository)

	tokenRepo := redisRepository.NewRedisRepo(redis.TokensClient)
	tokenService := jwt.NewJWTService(cfg, tokenRepo, auth.MaxAgeCookie*time.Duration(cfg.RefreshMaxAge))

	tagRepository := tag.NewTagRepo(postgres.Db)
	tagService := services.NewTagService(tagRepository)
	tags.NewTagsController(r, tagService, tokenService)

	auth.NewAuthController(r, authService, tokenService, cfg)
	users.NewUsersController(r, userService, tokenService)

	log.Info("Server starting...")

	if err := g.Run(cfg.HTTPServer.Address); err != nil {
		log.Fatal(err)
	}
}
