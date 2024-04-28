package rest_api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/services/auth_service"
	"task-manager/internal/domain/services/dto"
)

type IAuthController interface {
	Register(router *gin.Context)
	Login(router *gin.Context)
	Auth(router *gin.Context)
	Logout(router *gin.Context)
}

type AuthController struct {
	service auth_service.IAuthService
}

func NewUserController(gin *gin.Engine, service auth_service.IAuthService) *AuthController {
	controller := &AuthController{
		service: service,
	}

	r := gin.Group("auth")
	{
		r.POST("/register", controller.Register)
		r.POST("/login", controller.Login)
	}

	return controller
}

func (a AuthController) Register(ctx *gin.Context) {
	var registerUser dto.RegisterDTO

	err := ctx.ShouldBindJSON(&registerUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	regErr := a.service.Register(&entities.User{
		Email:     registerUser.Email,
		Username:  registerUser.Username,
		FirstName: registerUser.FirstName,
		LastName:  registerUser.LastName,
		Password:  registerUser.Password,
		Sphere:    registerUser.Sphere,
		Avatar:    registerUser.Avatar,
	})

	if regErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	//	validate payload
	//	catch errors
}

func (a AuthController) Login(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a AuthController) Auth(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a AuthController) Logout(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
