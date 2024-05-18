package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/services"
	"task-manager/internal/pkg"
)

type userService interface {
	GetUsers(queries *services.Queries) ([]entities.User, error)
	GetUserById(id string) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	CreateUser(user *entities.User) (id string, err error)
	UpdateUser(user *entities.User) error
	DeleteUser(id string) error
}

type tokenService interface {
	GenerateAccessToken(user *entities.User) (string, error)
	ParseAccessToken(tokenString string) (jwt.MapClaims, error)
	GenerateRefreshToken(id string) (string, error)
	CheckRefreshToken(refreshTokenString string) (string, error)
	DeleteRefreshToken(refreshTokenString string) error
}

type UsersController struct {
	userService  userService
	tokenService tokenService
}

func NewUsersController(gin *gin.RouterGroup, userService userService, tokenService tokenService) *UsersController {
	controller := &UsersController{
		userService:  userService,
		tokenService: tokenService,
	}

	r := gin.Group("/users")
	{
		r.GET("", controller.GetUsers)
		r.GET("/:id", controller.GetUserById)
		r.POST("", controller.CreateUser)
		r.PATCH("/:id", controller.UpdateUser)
		r.GET("/me", controller.GetMe)
		r.DELETE("/:id", controller.DeleteUser)
	}

	return controller
}

func (u UsersController) GetUsers(ctx *gin.Context) {
	search := ctx.Request.URL.Query().Get("search")
	page := ctx.Request.URL.Query().Get("page")

	users, err := u.userService.GetUsers(&services.Queries{
		Page:   page,
		Search: search,
	})

	if err != nil {
		logrus.Error("GET /users: | get users error: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(users)

}

func (u UsersController) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("id")
	user, err := u.userService.GetUserById(userId)

	if err != nil {
		logrus.Error("GET /users/:id: | user service error", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Something went wrong")
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (u UsersController) GetMe(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		logrus.Error("GET /users/me: | token is not provided")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Authorization header not provided")
		return
	}

	claims, err := u.tokenService.ParseAccessToken(token)
	if err != nil {
		logrus.Error("GET /users/me | token validation error", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Authorization header not provided")
		return
	}

	id := claims["id"].(string)

	user, err := u.userService.GetUserById(id)
	if err != nil {
		logrus.Error("GET /users/me: | user service error", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Something went wrong")
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (u UsersController) CreateUser(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u UsersController) UpdateUser(ctx *gin.Context) {
	var userDTO entities.User
	validate := validator.New()

	if err := ctx.ShouldBindJSON(&userDTO); err != nil {

		logrus.Error("PATCH /users/:id: | body parsing error", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := validate.Struct(&userDTO); err != nil {
		errorMessage := pkg.ExtractValidationErrors(err)

		logrus.Error("PATCH /users/:id: | body validation error", errorMessage)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errorMessage})
	}

	if err := u.userService.UpdateUser(&userDTO); err != nil {
		logrus.Error("PATCH /users/:id: | cannot update user: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Something went wrong")
	}

	ctx.JSON(http.StatusNoContent, userDTO)
}

func (u UsersController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := u.userService.DeleteUser(id); err != nil {
		logrus.Error("DELETE /users/:id: | body validation error", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
