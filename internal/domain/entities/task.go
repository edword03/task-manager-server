package entities

import "time"

type DueTime struct {
	From time.Time
	To   time.Time
}

type Task struct {
	ID         [16]byte
	Title      string
	Content    string
	CreateTime time.Time
	UpdateTime time.Time
	DueTime
	Priority  int
	Tags      []Tag
	Assignees User
}
