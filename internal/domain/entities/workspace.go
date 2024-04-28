package entities

import "time"

type Workspace struct {
	ID          [16]byte
	Name        string
	Type        string
	Description string
	Members     []User
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Tasks       []Task
}
