package DTO

import "task-manager/internal/domain/services/dto"

type UserDTO struct {
	Username  string `validate:"omitempty,alphanum" json:"username"`
	FirstName string `validate:"omitempty,alpha" json:"first_name"`
	LastName  string `validate:"omitempty,alpha" json:"last_name"`
	Sphere    string `validate:"alpha,omitempty" json:"sphere"`
	Avatar    string `json:"avatar"`
}

func ToDomainUserDTO(body UserDTO) *dto.UserDTO {
	return &dto.UserDTO{
		Username:  body.Username,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Sphere:    body.Sphere,
		Avatar:    body.Avatar,
	}
}
