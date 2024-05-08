package services

import (
	"errors"
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/services/dto"
	_ "task-manager/internal/domain/services/dto"
)

type AuthService struct {
	userRepo userRepository
}

func NewAuthService(repository userRepository) *AuthService {
	return &AuthService{userRepo: repository}
}

func (a AuthService) Register(payload *dto.RegisterDTO) (*entities.User, error) {
	existUser, err := a.userRepo.FindByEmail(payload.Email)

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

	var newUser = entities.User{
		Email:     payload.Email,
		Username:  payload.Username,
		FirstName: firstName,
		LastName:  lastName,
		Password:  payload.Password,
		Sphere:    payload.Sphere,
		Avatar:    avatar,
	}

	createdUser, err := a.userRepo.Create(&newUser)

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (a AuthService) Login(payload *dto.LoginDTO) (*entities.User, error) {
	existUser, err := a.userRepo.FindByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if existUser == nil {
		return nil, errors.New("user not found")
	}

	_, err = a.userRepo.ComparePassword(existUser.Password, payload.Password)
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return existUser, nil
}

func (a AuthService) Authenticate(id string) (*entities.User, error) {
	user, err := a.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
