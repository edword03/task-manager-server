package services

import (
	"errors"
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

func (u UserService) GetUserById(id string) (*entities.User, error) {
	user, err := u.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserService) GetUserByEmail(email string) (*entities.User, error) {
	user, err := u.userRepo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	if user == nil || user.Email == "" {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (u UserService) CreateUser(user *entities.User) (string, error) {
	user, err := u.userRepo.Create(user)

	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func (u UserService) UpdateUser(user *entities.User) error {
	err := u.userRepo.Update(user)

	if err != nil {
		return err
	}

	return nil
}

func (u UserService) DeleteUser(id string) error {
	user, err := u.userRepo.FindById(id)

	if err != nil {
		return err
	}

	if user.ID == "" {
		return errors.New("user not found")
	}

	if err := u.userRepo.Delete(id); err != nil {
		return err
	}

	return nil
}
