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
	"task-manager/internal/infrastructure/controllers/rest-api/auth/request"
	"task-manager/internal/infrastructure/controllers/rest-api/auth/request/mapper"
	"task-manager/internal/infrastructure/controllers/rest-api/auth/response"
	"time"
)

type IAuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Authenticate(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type Controller struct {
	authService  auth_service.IAuthService
	tokenService IJWTService
	log          *logrus.Logger
	cfg          *config.AppConfig
}

func NewAuthController(gin *gin.Engine, authService auth_service.IAuthService,
	tokenService IJWTService, cfg *config.AppConfig) *Controller {
	controller := &Controller{
		authService:  authService,
		cfg:          cfg,
		tokenService: tokenService,
	}

	r := gin.Group("auth")
	{
		r.POST("/register", controller.Register)
		r.POST("/login", controller.Login)
		r.POST("/authenticate", controller.Authenticate)
		r.POST("/logout", CheckTokenMiddleware(tokenService), controller.Logout)
	}

	return controller
}

func (a Controller) Register(ctx *gin.Context) {
	validate := validator.New()
	var registerUser request.RegisterDTO

	err := ctx.ShouldBindJSON(&registerUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

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

	newUser, regErr := a.authService.Register(mapper.ToDomainRegisterDTO(registerUser))

	if regErr != nil {
		logrus.Warnf("register service error: %v", regErr)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": regErr.Error()})
		return
	}

	if newUser == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user register service error"})
		return
	}

	accessToken, tokenErr := a.tokenService.GenerateAccessToken(newUser)

	if tokenErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": tokenErr.Error()})
		return
	}

	refreshToken, err := a.tokenService.GenerateRefreshToken(newUser.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("t", refreshToken, int(maxAgeCookie*time.Duration(a.cfg.RefreshMaxAge)), "/", "", false, true)

	ctx.JSON(http.StatusCreated, gin.H{"id": newUser.ID, "token": accessToken})
}

func (a Controller) Login(ctx *gin.Context) {
	validate := validator.New()
	var loginUser request.LoginDTO
	err := ctx.ShouldBindJSON(&loginUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logrus.Warnf("login request body error: %v, %v", err, loginUser)

		return
	}

	if validErr := validate.Struct(&loginUser); validErr != nil {
		var errMessage []string
		errors := validErr.(validator.ValidationErrors)
		for _, err := range errors {
			if err.ActualTag() == "required" {
				errMessage = append(errMessage, fmt.Sprintf("field %v is required", err.Field()))
			} else {
				errMessage = append(errMessage, fmt.Sprintf("field %v is not valid", err.Field()))
			}
		}

		logrus.Warnf("login request body error: %v, %v", strings.Join(errMessage, ", "), loginUser)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		return
	}

	user, err := a.authService.Login(mapper.ToDomainLoginDTO(loginUser))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := a.tokenService.GenerateAccessToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := a.tokenService.GenerateRefreshToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("t", refreshToken, int(maxAgeCookie*time.Duration(a.cfg.RefreshMaxAge)), "/", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{"user": response.ToUserResp(user), "token": accessToken})
}

func (a Controller) Authenticate(ctx *gin.Context) {
	token, err := ctx.Request.Cookie("t")

	if err != nil {
		logrus.Error("POST /authenticate | cookie is not passed: ", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if len(token.Value) == 0 {
		logrus.Error("POST /authenticate | token error: ", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token required"})
		return
	}

	userId, err := a.tokenService.CheckRefreshToken(token.Value)

	if err != nil {
		logrus.Error("POST /authenticate | check refresh token error: ", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is invalid or not exists"})
		return
	}

	if userId == "" {
		logrus.Error("POST /authenticate | userId empty: ", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token invalid"})
		return
	}

	if err := a.tokenService.DeleteRefreshToken(token.Value); err != nil {
		logrus.Error("POST /authenticate | delete refresh token error: ", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token invalid"})
		return
	}

	user, err := a.authService.Authenticate(userId)
	if err != nil {
		logrus.Error("POST /authenticate | token error: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := a.tokenService.GenerateAccessToken(user)
	if err != nil {
		logrus.Error("POST /authenticate | GenerateAccessToken error: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := a.tokenService.GenerateRefreshToken(userId)
	if err != nil {
		logrus.Error("POST /authenticate | GenerateRefreshToken error: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("t", refreshToken, int(maxAgeCookie*time.Duration(a.cfg.RefreshMaxAge)), "/", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"token": accessToken,
	})
}

func (a Controller) Logout(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("t")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	err = a.tokenService.DeleteRefreshToken(cookie.Value)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("t", "", -1, "/", "", false, true)
	ctx.JSON(http.StatusNoContent, gin.H{})
	return
}
