package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Email     string    `gorm:"type:varchar(255);unique;not null"`
	Username  string    `gorm:"type:varchar(255);unique;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Sphere    string    `gorm:"type:varchar(255);not null"`
	FirstName string
	LastName  string
	Avatar    string
}

func (u *User) BeforeCreate(db *gorm.DB) {
	u.ID = uuid.New()
}
