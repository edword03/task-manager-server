package DTO

import (
	"task-manager/internal/domain/services/dto"
)

type RegisterDTO struct {
	Email     string `validate:"required,email" json:"email"`
	Username  string `validate:"required,alphanum" json:"username"`
	FirstName string `validate:"omitempty,alpha" json:"first_name"`
	LastName  string `validate:"omitempty,alpha" json:"last_name"`
	Password  string `validate:"required,min=8" json:"password"`
	Sphere    string `validate:"required" json:"sphere"`
	Avatar    string `json:"avatar"`
}

type LoginDTO struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func ToDomainRegisterDTO(body RegisterDTO) *dto.RegisterDTO {
	return &dto.RegisterDTO{
		Email:     body.Email,
		Username:  body.Username,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Password:  body.Password,
		Sphere:    body.Sphere,
		Avatar:    body.Avatar,
	}
}

func ToDomainLoginDTO(body LoginDTO) *dto.LoginDTO {
	return &dto.LoginDTO{
		Email:    body.Email,
		Password: body.Password,
	}
}
