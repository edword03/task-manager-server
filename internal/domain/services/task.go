package services

import (
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/services/dto"
)

type taskRepository interface {
	GetById(id string) (*entities.Task, error)
	GetAll(page, pageSize int, searchTerm string) ([]*entities.Task, error)
	Create(task *entities.Task) (*entities.Task, error)
	Update(id string, task *dto.TaskDTO) error
	Delete(id string) error
}

type TaskService struct {
	taskRepository taskRepository
}

func NewTaskService(repo taskRepository) *TaskService {
	return &TaskService{repo}
}

func (t TaskService) GetTaskById(id string) (*entities.Task, error) {
	task, err := t.taskRepository.GetById(id)

	if err != nil {
		return nil, err
	}

	if task.ID == "" {
		return nil, NotFoundTaskError
	}

	return task, nil
}

func (t TaskService) GetTasks(page, pageSize int, searchTerm string) ([]*entities.Task, error) {
	tasks, err := t.taskRepository.GetAll(page, pageSize, searchTerm)

	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, NotFoundTaskError
	}

	return tasks, nil
}

func (t TaskService) Create(workspaceId, taskId string, task *dto.TaskDTO) (*entities.Task, error) {
	var (
		Content  = ""
		Priority = 3
		Assignee = entities.User{}
		DueTime  = entities.DueTime{
			From: nil,
			To:   nil,
		}
	)

	if task.Content != "" {
		Content = task.Content
	}

	if task.Priority != 0 {
		Priority = task.Priority
	}

	if task.Assignee.ID != "" {
		Assignee = task.Assignee
	}

	if !task.DueTime.From.IsZero() {
		DueTime.From = task.DueTime.From
	}

	if !task.DueTime.To.IsZero() {
		DueTime.To = task.DueTime.To
	}

	newTask := &entities.Task{
		TaskID:  taskId,
		Title:   task.Title,
		Content: Content,
		DueTime: entities.DueTime{
			From: DueTime.From,
			To:   DueTime.To,
		},
		Priority:    Priority,
		Tags:        task.Tags,
		Author:      task.Author,
		Assignee:    Assignee,
		WorkspaceId: workspaceId,
	}

	newTask, err := t.taskRepository.Create(newTask)

	if err != nil {
		return nil, CannotCreateTaskError
	}

	return newTask, nil
}

func (t TaskService) UpdateTask(workspaceId, id string, task *dto.TaskDTO) error {
	currentTask, err := t.taskRepository.GetById(id)

	if err != nil {
		return err
	}

	if currentTask.ID == "" {
		return NotFoundTaskError
	}

	if err := t.taskRepository.Update(id, task); err != nil {
		return CannotUpdateTaskError
	}

	return nil
}

func (t TaskService) DeleteTask(id string) error {
	task, err := t.taskRepository.GetById(id)

	if err != nil {
		return err
	}

	if task.ID == "" {
		return NotFoundTaskError
	}

	if err := t.taskRepository.Delete(id); err != nil {
		return CannotDeleteTaskError
	}

	return nil
}
