package user

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Username  string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Sphere    string `gorm:"type:varchar(255);not null"`
	FirstName string `gorm:"type:varchar(255);default:null"`
	LastName  string `gorm:"type:varchar(255);default:null"`
	Avatar    string `gorm:"type:varchar(255);default:null"`
}

func (u *User) BeforeCreate(*gorm.DB) error {
	u.ID = uuid.New().String()

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	u.Password = string(hashPass)
	return nil
}
