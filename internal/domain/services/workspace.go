package services

import (
	"github.com/sirupsen/logrus"
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/services/dto"
)

type workspaceRepository interface {
	Create(workspace *entities.Workspace, owner *entities.User) (*entities.Workspace, error)
	FindById(id string) (*entities.Workspace, error)
	FindAll(page int, pageSize int, searchTerm string) ([]*entities.Workspace, error)
	Update(workspaceID string, workspace *dto.WorkspaceDTO) error
	Delete(workspaceID string) error
}

type WorkspaceService struct {
	workspaceRepo workspaceRepository
}

func NewWorkspaceService(workspaceRepo workspaceRepository) *WorkspaceService {
	return &WorkspaceService{workspaceRepo: workspaceRepo}
}

func (s WorkspaceService) CreateWorkspace(workspace *dto.WorkspaceDTO, owner *entities.User) (*entities.Workspace, error) {
	var (
		Description = ""
		Logo        = ""
	)

	if workspace.Description != "" {
		Description = workspace.Description
	}

	if workspace.Logo != "" {
		Logo = workspace.Logo
	}

	newWorkspace := &entities.Workspace{
		Name:        workspace.Name,
		Type:        workspace.Type,
		Description: Description,
		Logo:        Logo,
	}

	newWorkspace, err := s.workspaceRepo.Create(newWorkspace, owner)
	if err != nil {
		logrus.Error("workspace service: cannot create - ", err)
		return nil, CannotCreateWorkspaceError
	}

	return newWorkspace, nil
}

func (s WorkspaceService) GetWorkspaceById(id string) (*entities.Workspace, error) {
	workspace, err := s.workspaceRepo.FindById(id)
	if err != nil {
		return nil, NotFoundWorkspaceError
	}

	return workspace, nil
}

func (s WorkspaceService) GetAllWorkspaces(page int, pageSize int, searchTerm string) ([]*entities.Workspace, error) {
	workspaces, err := s.workspaceRepo.FindAll(page, pageSize, searchTerm)

	if err != nil {
		return nil, err
	}

	if len(workspaces) == 0 {
		return nil, NotFoundWorkspaceError
	}

	return workspaces, nil
}

func (s WorkspaceService) UpdateWorkspace(workspaceID string, workspace *dto.WorkspaceDTO) error {
	currentWorkspace, err := s.workspaceRepo.FindById(workspaceID)

	if err != nil {
		return NotFoundWorkspaceError
	}

	if currentWorkspace.ID == "" {
		return NotFoundWorkspaceError
	}

	if err := s.workspaceRepo.Update(workspaceID, workspace); err != nil {
		return CannotUpdateWorkspaceError
	}

	return nil
}

func (s WorkspaceService) DeleteWorkspace(workspaceID string) error {
	workspace, err := s.workspaceRepo.FindById(workspaceID)

	if err != nil {
		return NotFoundWorkspaceError
	}

	if workspace.ID == "" {
		return NotFoundWorkspaceError
	}

	if err := s.workspaceRepo.Delete(workspaceID); err != nil {
		return CannotDeleteWorkspaceError
	}

	return nil
}
