package entities

import "time"

type Settings struct {
}

type Workspace struct {
	ID          [16]byte
	Name        string
	Type        string
	Description string
	Logo        string
	Members     []User
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Tasks       []Task
	Labels      []Label
	Roles       []Role
	Status      string
	Settings    map[string]interface{}
	Features    []Feature
}
