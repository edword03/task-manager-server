package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"task-manager/internal/domain/services/auth_service"
	"task-manager/internal/infrastructure/config"
	"task-manager/internal/infrastructure/controllers/rest-api/auth/jwt"
	"task-manager/internal/infrastructure/controllers/rest-api/auth/request"
	"task-manager/internal/infrastructure/controllers/rest-api/auth/request/mapper"
	"task-manager/internal/infrastructure/database/redis"
	redisRepo "task-manager/internal/infrastructure/database/redis/repositories"
	"time"
)

type IAuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Auth(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type Controller struct {
	authService  auth_service.IAuthService
	tokenService jwt.IJWTService
	tokenRepo    redisRepo.ITokenRepository
	log          *logrus.Logger
	cfg          *config.AppConfig
}

const (
	maxAgeCookie = time.Hour * 24
)

func NewUserController(gin *gin.Engine, authService auth_service.IAuthService,
	tokenService jwt.IJWTService, cfg *config.AppConfig) *Controller {
	controller := &Controller{
		authService:  authService,
		cfg:          cfg,
		tokenService: tokenService,
		tokenRepo:    redisRepo.NewTokenRepo(redis.TokensClient),
	}

	r := gin.Group("auth")
	{
		r.POST("/register", controller.Register)
		r.POST("/login", controller.Login)
		r.POST("/logout", controller.Logout)
	}

	return controller
}

func (a Controller) Register(ctx *gin.Context) {
	validate := validator.New()
	var registerUser request.RegisterDTO

	err := ctx.ShouldBindJSON(&registerUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		logrus.Warnf("register request body error: %v, %v", err, registerUser)

		return
	}

	if validErr := validate.Struct(&registerUser); validErr != nil {
		var errMessage []string
		errors := validErr.(validator.ValidationErrors)

		for _, err := range errors {
			if err.ActualTag() == "required" {
				errMessage = append(errMessage, fmt.Sprintf("field %v is required", err.Field()))
			} else {
				errMessage = append(errMessage, fmt.Sprintf("field %v is not valid", err.Field()))
			}
		}

		logrus.Warnf("register request body error: %v, %v", strings.Join(errMessage, ", "), registerUser)

		ctx.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		return
	}

	newUser, regErr := a.authService.Register(mapper.ToDomainDTO(registerUser))

	if regErr != nil {
		logrus.Warnf("register service error: %v", regErr)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": regErr.Error()})
		return
	}

	if newUser == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user register service error"})
	}

	accessToken, tokenErr := a.tokenService.GenerateAccessToken(newUser)
	refreshToken := a.tokenService.GenerateRefreshToken()

	if tokenErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": tokenErr.Error()})
		return
	}

	ctx.SetCookie("t", refreshToken, int(maxAgeCookie*time.Duration(a.cfg.RefreshMaxAge)), "/", "", false, true)

	tokenRepoErr := a.tokenRepo.Set(refreshToken, newUser.ID, maxAgeCookie*time.Duration(a.cfg.RefreshMaxAge))
	if tokenRepoErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": tokenRepoErr.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": newUser.ID, "token": accessToken})
}

func (a Controller) Login(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a Controller) Auth(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a Controller) Logout(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
