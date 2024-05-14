package services

import (
	"task-manager/internal/domain/entities"
)

type Queries struct {
	Page   string
	Search string
}

type UserService struct {
	userRepo userRepository
}

func NewUserService(repository userRepository) *UserService {
	return &UserService{
		userRepo: repository,
	}
}

func (u UserService) GetUsers(queries *Queries) ([]entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetUserById(id string) (entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) CreateUser(user *entities.User) (id string, err error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) UpdateUser(user *entities.User) (entity entities.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) DeleteUser(id string) error {
	//TODO implement me
	panic("implement me")
}
