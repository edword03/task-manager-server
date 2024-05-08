package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/services"
)

//type IUsersController interface {
//	GetUsers(w http.ResponseWriter, r *http.Request)
//	GetUserById(ctx *gin.Context)
//	CreateUser(ctx *gin.Context)
//	UpdateUser(ctx *gin.Context)
//	DeleteUser(ctx *gin.Context)
//}

type userService interface {
	GetUsers(queries *services.Queries) ([]entities.User, error)
	GetUserById(id string) (entities.User, error)
	CreateUser(user *entities.User) (id string, err error)
	UpdateUser(user *entities.User) (entity entities.User, err error)
	DeleteUser(id string) error
}

type UsersController struct {
	userService userService
}

func NewUsersController(gin *gin.Engine, userService userService) *UsersController {
	controller := &UsersController{
		userService: userService,
	}

	r := gin.Group("/users")
	{
		r.GET("", controller.GetUsers)
		r.GET("/:id", controller.GetUserById)
		r.POST("", controller.CreateUser)
		r.PUT("/:id", controller.UpdateUser)

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
		logrus.Error("")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(users)

}

func (u UsersController) GetUserById(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u UsersController) CreateUser(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u UsersController) UpdateUser(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u UsersController) DeleteUser(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
