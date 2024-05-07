package user_service

import (
	"task-manager/internal/domain/entities"
	"task-manager/internal/infrastructure/database/postgres/repositories"
)

type IUserService interface {
	GetUsers(queries *Queries) ([]entities.User, error)
	GetUserById(id string) (entities.User, error)
	CreateUser(user *entities.User) (id string, err error)
	UpdateUser(user *entities.User) (entity entities.User, err error)
	DeleteUser(id string) error
}

type Queries struct {
	Page   string
	Search string
}

type UserService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

//func (u UserService) GetUsers(queries Queries) ([]entities.User, error) {
//	user, err := u.repository.FindAll(query)
//}

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
