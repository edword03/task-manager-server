package workspaces

import "task-manager/internal/domain/services/dto"

type WorkspaceDTO struct {
	Name        string `validate:"required,alphanum" json:"name"`
	Type        string `validate:"required,alphanum" json:"type"`
	Description string `validate:"omitempty,alphanum" json:"description"`
	OwnerId     string `validate:"required,uuid" json:"owner_id"`
}

func ToDomainWorkspaceDTO(body WorkspaceDTO) *dto.WorkspaceDTO {
	return &dto.WorkspaceDTO{
		Name:        body.Name,
		Type:        body.Type,
		Description: body.Description,
	}
}
