package mappers

import (
	"task-manager/internal/domain/entities"
	"task-manager/internal/infrastructure/database/postgres/model"
)

func ToDBUser(user *entities.User) *model.User {
	return &model.User{
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Sphere:    user.Sphere,
		Avatar:    user.Avatar,
	}

}

func ToDomainUser(user *model.User) *entities.User {
	return &entities.User{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Sphere:    user.Sphere,
		Avatar:    user.Avatar,
	}
}
