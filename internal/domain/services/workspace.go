package services

import "task-manager/internal/domain/entities"

type workspaceRepository interface {
	Create(workspace entities.Workspace) error
	FindById(id string) (*entities.Workspace, error)
	FindAll() ([]entities.Workspace, error)
	GetUserWorkspaces(userID string) ([]entities.Workspace, error)
	Update(workspaceID string, workspace entities.Workspace) error
	Delete(workspaceID string) error
}

type WorkspaceService struct {
	workspaceRepo workspaceRepository
}

func NewWorkspaceService(workspaceRepo workspaceRepository) *WorkspaceService {
	return &WorkspaceService{workspaceRepo: workspaceRepo}
}
