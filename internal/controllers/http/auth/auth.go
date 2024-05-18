package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"task-manager/internal/config"
	"task-manager/internal/controllers/http/DTO"
	"task-manager/internal/controllers/http/response"
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/services/dto"
	"task-manager/internal/pkg"
	"time"
)

type authService interface {
	Register(payload *dto.RegisterDTO) (*entities.User, error)
	Login(payload *dto.LoginDTO) (*entities.User, error)
	Authenticate(id string) (*entities.User, error)
}

type jwtService interface {
	GenerateAccessToken(user *entities.User) (string, error)
	ParseAccessToken(tokenString string) (jwt.MapClaims, error)
	GenerateRefreshToken(id string) (string, error)
	CheckRefreshToken(refreshTokenString string) (string, error)
	DeleteRefreshToken(refreshTokenString string) error
}

type Controller struct {
	authService  authService
	tokenService jwtService
	log          *logrus.Logger
	cfg          *config.AppConfig
}

func NewAuthController(gin *gin.RouterGroup, authService authService,
	tokenService jwtService, cfg *config.AppConfig) *Controller {
	controller := &Controller{
		authService:  authService,
		cfg:          cfg,
		tokenService: tokenService,
	}

	gin.Group("auth")
	{
		gin.POST("/register", controller.Register)
		gin.POST("/login", controller.Login)
		gin.POST("/authenticate", controller.Authenticate)
		gin.POST("/logout", CheckTokenMiddleware(tokenService), controller.Logout)
	}

	return controller
}

func (a Controller) Register(ctx *gin.Context) {
	validate := validator.New()
	var registerUser DTO.RegisterDTO

	if err := ctx.ShouldBindJSON(&registerUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		logrus.Warnf("register request body error: %v, %v", err, registerUser)

		return
	}

	if validErr := validate.Struct(&registerUser); validErr != nil {
		errorMessage := pkg.ExtractValidationErrors(validErr)

		logrus.Warnf("register request body error: %v, %v", strings.Join(errorMessage, ", "), registerUser)

		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	newUser, regErr := a.authService.Register(DTO.ToDomainRegisterDTO(registerUser))

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

	ctx.SetCookie("t", refreshToken, int(MaxAgeCookie*time.Duration(a.cfg.RefreshMaxAge)), "/", "", false, true)

	ctx.JSON(http.StatusCreated, gin.H{"id": newUser.ID, "token": accessToken})
}

func (a Controller) Login(ctx *gin.Context) {
	validate := validator.New()
	var loginUser DTO.LoginDTO

	if err := ctx.ShouldBindJSON(&loginUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logrus.Warnf("login request body error: %v, %v", err, loginUser)

		return
	}

	if validErr := validate.Struct(&loginUser); validErr != nil {
		errorMessage := pkg.ExtractValidationErrors(validErr)

		logrus.Warnf("login request body error: %v, %v", strings.Join(errorMessage, ", "), loginUser)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	user, err := a.authService.Login(DTO.ToDomainLoginDTO(loginUser))
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

	ctx.SetCookie("t", refreshToken, int(MaxAgeCookie*time.Duration(a.cfg.RefreshMaxAge)), "/", "", false, true)

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

	ctx.SetCookie("t", refreshToken, int(MaxAgeCookie*time.Duration(a.cfg.RefreshMaxAge)), "/", "", false, true)

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
