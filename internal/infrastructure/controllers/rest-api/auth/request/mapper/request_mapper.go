package mapper

import (
	"task-manager/internal/domain/services/dto"
	"task-manager/internal/infrastructure/controllers/rest-api/auth/request"
)

func ToDomainDTO(body request.RegisterDTO) *dto.RegisterDTO {
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
