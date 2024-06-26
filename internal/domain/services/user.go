package services

import (
	"errors"
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/services/dto"
)

type UserService struct {
	userRepo userRepository
}

func NewUserService(repository userRepository) *UserService {
	return &UserService{
		userRepo: repository,
	}
}

func (u UserService) GetUsers(page, pageSize int, searchTerm string) ([]*entities.User, error) {
	users, err := u.userRepo.FindAll(page, pageSize, searchTerm)

	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("users not found")
	}

	return users, nil
}

func (u UserService) GetUserById(id string) (*entities.User, error) {
	user, err := u.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	if user.ID == "" {
		return nil, errors.New("user not found")
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

func (u UserService) CreateUser(payload *dto.RegisterDTO) (*entities.User, error) {
	existUser, err := u.userRepo.FindByEmail(payload.Email)

	if err != nil {
		return nil, err
	}

	if existUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	var avatar = "/mock-avatar.png"
	var firstName string
	var lastName string

	if payload.Avatar != "" {
		avatar = payload.Avatar
	}

	if payload.FirstName != "" {
		firstName = payload.FirstName
	}

	if payload.LastName != "" {
		lastName = payload.LastName
	}

	newUser, err := u.userRepo.Create(&entities.User{
		Email:     payload.Email,
		Username:  payload.Username,
		FirstName: firstName,
		LastName:  lastName,
		Password:  payload.Password,
		Sphere:    payload.Sphere,
		Avatar:    avatar,
	})

	return newUser, nil
}

func (u UserService) UpdateUser(userId string, user *dto.UserDTO) error {
	currentUser, err := u.userRepo.FindById(userId)

	if err != nil {
		return err
	}

	if currentUser == nil {
		return errors.New("user not found")
	}

	err = u.userRepo.Update(userId, user)

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
