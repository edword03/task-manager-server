package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	TaskID      string `gorm:"unique;not null; varchar(20)"`
	WorkspaceID string
	Title       string `gorm:"not null; varchar(255)"`
	Content     string `gorm:"varchar(255)"`
	FromTime    *time.Time
	ToTime      *time.Time
	Priority    int `gorm:"not null"`
	AuthorID    string
	AssigneeID  string
	Author      User  `gorm:"one2many:user;foreignkey:AuthorID"`
	Assignee    User  `gorm:"one2many:user;foreignkey:AssigneeID"`
	Tags        []Tag `gorm:"many2many:task_tags;"`
}

func (task *Task) BeforeCreate(*gorm.DB) error {
	task.ID = uuid.New().String()

	return nil
}
