package response

import (
	"task-manager/internal/domain/entities"
	"time"
)

type TaskResponse struct {
	ID          string         `json:"id"`
	TaskID      string         `json:"task_id"`
	WorkspaceId string         `json:"workspace_id"`
	Title       string         `json:"title"`
	Content     string         `json:"content"`
	FromTime    time.Time      `json:"from_time"`
	ToTime      time.Time      `json:"to_time"`
	Priority    int            `json:"priority"`
	Author      entities.User  `json:"author"`
	Assignee    entities.User  `json:"assignee"`
	Tags        []entities.Tag `json:"tags"`
}

func ToTaskResponse(task *entities.Task) *TaskResponse {
	return &TaskResponse{
		ID:          task.ID,
		TaskID:      task.TaskID,
		WorkspaceId: task.WorkspaceId,
		Title:       task.Title,
		Content:     task.Content,
		FromTime:    task.DueTime.From,
		ToTime:      task.DueTime.To,
		Priority:    task.Priority,
		Author:      task.Author,
		Assignee:    task.Assignee,
		Tags:        task.Tags,
	}
}
