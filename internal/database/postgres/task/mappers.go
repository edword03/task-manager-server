package task

import (
	"task-manager/internal/domain/entities"
)

func ToDBTask(task *entities.Task) *Task {
	return &Task{
		TaskID:      task.TaskID,
		WorkspaceId: task.WorkspaceId,
		Title:       task.Title,
		Content:     task.Content,
		FromTime:    task.DueTime.From,
		ToTime:      task.DueTime.To,
		Priority:    task.Priority,
	}
}

func ToDomainTask(task *Task) *entities.Task {
	return &entities.Task{
		ID:          task.ID,
		TaskID:      task.TaskID,
		WorkspaceId: task.WorkspaceId,
		Title:       task.Title,
		Content:     task.Content,
		CreateTime:  task.CreatedAt,
		UpdateTime:  task.UpdatedAt,
		DueTime: entities.DueTime{
			From: task.FromTime,
			To:   task.ToTime,
		},
		Priority: task.Priority,
	}
}

func ToDomainTasks(tasks []*Task) []*entities.Task {
	var results []*entities.Task

	for _, task := range tasks {
		results = append(results, ToDomainTask(task))
	}

	return results
}
