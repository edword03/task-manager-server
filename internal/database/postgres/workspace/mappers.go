package workspace

import (
	"task-manager/internal/database/postgres/models"
	"task-manager/internal/domain/entities"
)

func ToDBWorkspace(workspace *entities.Workspace, owner *models.User) *models.Workspace {
	return &models.Workspace{
		Name:        workspace.Name,
		Type:        workspace.Type,
		Description: workspace.Description,
		Logo:        workspace.Logo,
		OwnerID:     owner.ID,
		Status:      workspace.Status,
	}
}

func ToDomainWorkspace(workspace *models.Workspace) *entities.Workspace {
	return &entities.Workspace{
		ID:          workspace.ID,
		Name:        workspace.Name,
		Type:        workspace.Type,
		Description: workspace.Description,
		Logo:        workspace.Logo,
		CreatedAt:   workspace.CreatedAt,
		UpdatedAt:   workspace.UpdatedAt,
		Status:      workspace.Status,
	}
}

func ToDomainWorkspaces(workspaces []*models.Workspace) []*entities.Workspace {
	var result []*entities.Workspace

	for _, workspace := range workspaces {
		result = append(result, ToDomainWorkspace(workspace))
	}

	return result
}
