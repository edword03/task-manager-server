package entities

import "time"

type Settings struct {
}

type Workspace struct {
	ID          string
	Name        string
	Type        string
	Description string
	Logo        string
	Owner       User
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
