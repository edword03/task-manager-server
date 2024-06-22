package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Workspace struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Name        string `gorm:"varchar(255)"`
	Type        string `gorm:"varchar(255)"`
	Description string `gorm:"varchar(255)"`
	Logo        string `gorm:"varchar(255)"`
	OwnerID     string
	Owner       User   `gorm:"foreignKey:OwnerID"`
	Members     []User `gorm:"many2many:user_members;"`
	Tasks       []Task `gorm:"one2many:task;foreignKey:WorkspaceID"`
	Tags        []Tag  `gorm:"one2many:tag;foreignKey:WorkspaceID"`
	Status      string `gorm:"varchar(255)"`
}

func (workspace *Workspace) BeforeCreate(*gorm.DB) error {
	workspace.ID = uuid.New().String()

	return nil
}
