package user

import (
	"task-manager/internal/database/postgres/models"
	"task-manager/internal/domain/entities"
)

func ToDBUser(user *entities.User) *models.User {
	return &models.User{
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Sphere:    user.Sphere,
		Avatar:    user.Avatar,
	}

}

func ToDomainUser(user *models.User) *entities.User {
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

func ToDomainUsers(users []*models.User) []*entities.User {
	var result []*entities.User
	for _, user := range users {
		result = append(result, ToDomainUser(user))
	}

	return result
}
