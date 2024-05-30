package task

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"task-manager/internal/database/postgres/tag"
	"task-manager/internal/database/postgres/user"
	"time"
)

type Task struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	TaskID      string `gorm:"unique;not null; varchar(20)"`
	WorkspaceId string
	Title       string    `gorm:"not null; varchar(255)"`
	Content     string    `gorm:"varchar(255)"`
	FromTime    time.Time `gorm:"not null;type:datetime"`
	ToTime      time.Time `gorm:"not null;type:datetime"`
	Priority    int       `gorm:"not null"`
	Author      user.User `gorm:"one2many:user;"`
	Assignee    user.User `gorm:"one2many:user;"`
	Tags        []tag.Tag `gorm:"many2many:task_tags;"`
}

func (task *Task) BeforeCreate() error {
	task.ID = uuid.New().String()

	return nil
}
