package services

import (
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/services/dto"
)

type userRepository interface {
	Create(user *entities.User) (*entities.User, error)
	FindById(id string) (*entities.User, error)
	FindByUsername(username string) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	ComparePassword(password, passwordDto string) (bool, error)
	FindAll(page, pageSize int, searchTerm string) ([]*entities.User, error)
	Update(userId string, user *dto.UserDTO) error
	Delete(id string) error
}
