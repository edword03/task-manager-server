package dto

import (
	"task-manager/internal/domain/entities"
)

type TaskDTO struct {
	Title   string
	Content string
	entities.DueTime
	Priority    int
	Tags        []entities.Tag
	Author      entities.User
	Assignee    entities.User
	WorkspaceId string
}
