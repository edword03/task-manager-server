package user

import (
	"task-manager/internal/domain/entities"
)

func ToDBUser(user *entities.User) *User {
	return &User{
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Sphere:    user.Sphere,
		Avatar:    user.Avatar,
	}

}

func ToDomainUser(user *User) *entities.User {
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

func ToDomainUsers(users []*User) []*entities.User {
	var result []*entities.User
	for _, user := range users {
		result = append(result, ToDomainUser(user))
	}

	return result
}
