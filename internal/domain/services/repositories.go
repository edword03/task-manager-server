package services

import "task-manager/internal/domain/entities"

type userRepository interface {
	Create(user *entities.User) (*entities.User, error)
	FindById(id string) (*entities.User, error)
	FindByUsername(username string) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	ComparePassword(password, passwordDto string) (bool, error)
	FindAll(query string) ([]*entities.User, error)
	Update(user *entities.User) error
	Delete(id string) error
}
