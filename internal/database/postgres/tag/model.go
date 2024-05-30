package tag

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"task-manager/internal/database/postgres/task"
)

type Tag struct {
	gorm.Model
	ID    string      `gorm:"primaryKey"`
	Name  string      `gorm:"type:varchar(255);unique;not null"`
	Color string      `gorm:"type:varchar(255);default:null"`
	Tasks []task.Task `gorm:"many2many:tag_tasks;"`
}

func (u *Tag) BeforeCreate(*gorm.DB) error {
	u.ID = uuid.New().String()

	return nil
}
