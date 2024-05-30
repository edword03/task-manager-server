package entities

import "time"

type DueTime struct {
	From time.Time
	To   time.Time
}

type Task struct {
	ID          string
	TaskID      string
	WorkspaceId string
	Title       string
	Content     string
	CreateTime  time.Time
	UpdateTime  time.Time
	DueTime
	Priority int
	Tags     []Tag
	Author   User
	Assignee User
}
