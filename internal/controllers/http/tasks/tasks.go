package tasks

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"task-manager/internal/controllers/http/response"
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/services/dto"
)

type taskService interface {
	GetTaskById(id string) (*entities.Task, error)
	GetTasks(page, pageSize int, searchTerm string) ([]*entities.Task, error)
	Create(workspaceId, taskId string, task *dto.TaskDTO) (*entities.Task, error)
	UpdateTask(workspaceId, id string, task *dto.TaskDTO) error
	DeleteTask(id string) error
}

type Controller struct {
	taskService taskService
}

func NewTasksController(gin *gin.RouterGroup, taskService taskService) *Controller {
	controller := &Controller{
		taskService: taskService,
	}

	return controller
}

func (t Controller) GetTasks(ctx *gin.Context) {
	searchTerm := ctx.Request.URL.Query().Get("search")
	pageStr := ctx.Request.URL.Query().Get("page")
	limitStr := ctx.Request.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		logrus.Error("GET /users: | strconv.Atoi error: ", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		logrus.Error("GET /users: | strconv.Atoi error: ", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	tasks, err := t.taskService.GetTasks(page, limit, searchTerm)
	offset := (page - 1) * limit

	if err != nil {
		logrus.Error("GET ")
	}

	var tasksResponse []*response.TaskResponse

	for _, task := range tasks {
		tasksResponse = append(tasksResponse, response.ToTaskResponse(task))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"tasks":   tasks,
		"_offset": offset,
		"_limit":  limit,
		"_page":   page,
		"_total":  len(tasksResponse),
	})
}

func (t Controller) GetTaskById(ctx *gin.Context) {
	taskId := ctx.Param("id")

	task, err := t.taskService.GetTaskById(taskId)
	if err != nil {
		logrus.Error("GET /users: | GetTaskById error: ", err)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	ctx.JSON(http.StatusOK, response.ToTaskResponse(task))
}

func (t Controller) Create(ctx *gin.Context) {

}
