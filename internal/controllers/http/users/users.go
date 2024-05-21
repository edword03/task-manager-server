package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"task-manager/internal/controllers/http/DTO"
	"task-manager/internal/controllers/http/auth"
	"task-manager/internal/controllers/http/response"
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/services"
	"task-manager/internal/domain/services/dto"
	"task-manager/internal/pkg"
)

type userService interface {
	GetUsers(queries *services.Queries) ([]entities.User, error)
	GetUserById(id string) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	CreateUser(payload *dto.RegisterDTO) (*entities.User, error)
	UpdateUser(userId string, user *dto.UserDTO) error
	DeleteUser(id string) error
}

type tokenService interface {
	GenerateAccessToken(user *entities.User) (string, error)
	ParseAccessToken(tokenString string) (jwt.MapClaims, error)
	GenerateRefreshToken(id string) (string, error)
	CheckRefreshToken(refreshTokenString string) (string, error)
	DeleteRefreshToken(refreshTokenString string) error
}

type Controller struct {
	userService  userService
	tokenService tokenService
}

func NewUsersController(gin *gin.RouterGroup, userService userService, tokenService tokenService) *Controller {
	controller := &Controller{
		userService:  userService,
		tokenService: tokenService,
	}

	r := gin.Group("/users", auth.CheckTokenMiddleware(tokenService))
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

func (u Controller) GetUsers(ctx *gin.Context) {
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

func (u Controller) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("id")
	user, err := u.userService.GetUserById(userId)

	if err != nil {
		logrus.Error("GET /users/:id: | user service error", err)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.ToUserResp(user))
}

func (u Controller) GetMe(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	splitToken := strings.Split(token, "Bearer ")

	if token == "" {
		logrus.Error("GET /users/me: | token is not provided")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Authorization header not provided")
		return
	}

	claims, err := u.tokenService.ParseAccessToken(splitToken[1])
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

	ctx.JSON(http.StatusOK, response.ToUserResp(user))
}

func (u Controller) CreateUser(ctx *gin.Context) {
	var userDTO DTO.RegisterDTO
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

	id, err := u.userService.CreateUser(DTO.ToDomainRegisterDTO(userDTO))

	if err != nil {
		logrus.Error("PATCH /users/:id: | cannot update user: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Something went wrong")
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (u Controller) UpdateUser(ctx *gin.Context) {
	var userDTO DTO.UserDTO
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

	id := ctx.Param("id")

	if err := u.userService.UpdateUser(id, DTO.ToDomainUserDTO(userDTO)); err != nil {
		logrus.Error("PATCH /users/:id: | cannot update user: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Something went wrong")
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (u Controller) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := u.userService.DeleteUser(id); err != nil {
		logrus.Error("DELETE /users/:id: | body validation error", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
