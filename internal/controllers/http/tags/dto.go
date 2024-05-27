package tags

import "task-manager/internal/domain/services/dto"

type TagDTO struct {
	Name  string `validate:"required,alphanum" json:"name"`
	Color string `validate:"omitempty,iscolor" json:"color"`
}

func ToDomainTagDTO(body TagDTO) *dto.Tag {
	return &dto.Tag{
		Name:  body.Name,
		Color: body.Color,
	}
}
