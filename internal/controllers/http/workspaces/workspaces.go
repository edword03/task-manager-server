package workspaces

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/services"
	"task-manager/internal/domain/services/dto"
	"task-manager/internal/pkg"
)

type workspaceService interface {
	CreateWorkspace(workspace *dto.WorkspaceDTO, owner *entities.User) (*entities.Workspace, error)
	GetWorkspaceById(id string) (*entities.Workspace, error)
	GetAllWorkspaces(page int, pageSize int, searchTerm string) ([]*entities.Workspace, error)
	UpdateWorkspace(workspaceID string, workspace *dto.WorkspaceDTO) error
	DeleteWorkspace(workspaceID string) error
}

type userService interface {
	CreateUser(payload *dto.RegisterDTO) (*entities.User, error)
	GetUserById(id string) (*entities.User, error)
	GetUsers(page, pageSize int, searchTerm string) ([]*entities.User, error)
	UpdateUser(userId string, user *dto.UserDTO) error
	DeleteUser(id string) error
}

type Controller struct {
	workspaceService workspaceService
	userService      userService
}

func NewWorkspaceController(gin *gin.RouterGroup, workspaceService workspaceService, userService userService) *Controller {
	controller := &Controller{
		workspaceService: workspaceService,
		userService:      userService,
	}

	r := gin.Group("workspaces")
	{
		r.POST("", controller.CreateWorkspace)
		r.GET("/:id", controller.GetWorkspaceById)
	}

	return controller
}

func (c *Controller) CreateWorkspace(ctx *gin.Context) {
	var workspaceDTO WorkspaceDTO
	validate := validator.New()

	if err := ctx.ShouldBind(&workspaceDTO); err != nil {
		logrus.Error("POST /workspaces | body parsing error :", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := validate.Struct(workspaceDTO); err != nil {
		errorMessage := pkg.ExtractValidationErrors(err)

		logrus.Error("POST /workspaces | body validation error: ", errorMessage)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})

		return
	}

	owner, err := c.userService.GetUserById(workspaceDTO.OwnerId)

	if err != nil {
		logrus.Error("POST /workspaces | body parsing error :", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workspace, err := c.workspaceService.CreateWorkspace(ToDomainWorkspaceDTO(workspaceDTO), owner)

	if errors.Is(err, services.ExistWorkspaceError) {
		logrus.Error("POST /workspaces | workspace is exists :", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": services.ExistTagError.Error()})
		return
	}

	if err != nil {
		logrus.Error("POST /workspaces | cannot create workspace :", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	ctx.JSON(http.StatusCreated, workspace)
}

func (c *Controller) GetWorkspaceById(ctx *gin.Context) {
	id := ctx.Param("id")

	fmt.Printf(id)

	workspace, err := c.workspaceService.GetWorkspaceById(id)

	if err != nil {
		logrus.Error("POST /workspaces | workspace error: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": services.NotFoundWorkspaceError})
		return
	}

	fmt.Println(workspace)

	ctx.JSON(http.StatusOK, workspace)
}
