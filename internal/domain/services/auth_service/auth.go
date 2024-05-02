package auth_service

import (
	"errors"
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/repositories"
	"task-manager/internal/domain/services/dto"
	_ "task-manager/internal/domain/services/dto"
)

type IAuthService interface {
	Register(payload *dto.RegisterDTO) (*entities.User, error)
	Login(payload dto.LoginDTO) (*entities.User, error)
	Authenticate(id string) (*entities.User, error)
}

type AuthService struct {
	repository repositories.UserRepository
}

func NewAuthService(repository repositories.UserRepository) *AuthService {
	return &AuthService{repository: repository}
}

func (a AuthService) Register(payload *dto.RegisterDTO) (*entities.User, error) {
	existUser, err := a.repository.FindByEmail(payload.Email)

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

	createdUser, err := a.repository.Create(&newUser)

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (a AuthService) Login(payload dto.LoginDTO) (*entities.User, error) {
	existUser, err := a.repository.FindByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if existUser == nil {
		return nil, errors.New("user not exist")
	}

	isCorrectPass, passErr := a.repository.ComparePassword(existUser.Password, payload.Password)

	if passErr != nil {
		return nil, passErr
	}

	if !isCorrectPass {
		return nil, errors.New("incorrect password")
	}

	return existUser, nil
}

func (a AuthService) Authenticate(id string) (*entities.User, error) {
	user, err := a.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
