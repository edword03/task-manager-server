package repositories

import (
	"task-manager/internal/domain/entities"
)

type UserRepository interface {
	Create(user *entities.User) error
	FindById(id string) (*entities.User, error)
	FindByUsername(username string) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	ComparePassword(password, passwordDto string) (bool, error)
	FindAll() ([]*entities.User, error)
	Update(user *entities.User) error
	Delete(id string) error
}
