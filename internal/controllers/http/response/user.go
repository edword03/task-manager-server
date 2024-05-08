package response

import "task-manager/internal/domain/entities"

type UserResp struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Sphere    string `json:"sphere"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

func ToUserResp(user *entities.User) *UserResp {
	return &UserResp{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Sphere:    user.Sphere,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Avatar:    user.Avatar,
	}
}
