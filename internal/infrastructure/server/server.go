package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"task-manager/internal/domain/services/auth_service"
	"task-manager/internal/infrastructure/config"
	restapi "task-manager/internal/infrastructure/controllers/rest-api/auth"
	"task-manager/internal/infrastructure/controllers/rest-api/auth/jwt"
	"task-manager/internal/infrastructure/database/postgres"
	"task-manager/internal/infrastructure/database/postgres/repositories"
	"task-manager/internal/infrastructure/lib/logger"
)

func New(cfg *config.AppConfig) {
	r := gin.Default()

	logger.SetupLogger(cfg.Env)

	r.Use(gin.Logger())
	r.Use(cors.Default())
	r.Use(gin.Recovery())

	userRepository := repositories.NewUserRepo(postgres.Db)
	userService := auth_service.NewAuthService(userRepository)

	tokenService := jwt.NewJWTService(cfg)

	restapi.NewUserController(r, userService, tokenService, cfg)

	log.Info("Server starting...")

	if err := r.Run(cfg.HTTPServer.Address); err != nil {
		log.Fatal(err)
	}
}
