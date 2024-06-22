package workspace

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"task-manager/internal/database/postgres/models"
	"task-manager/internal/database/postgres/user"
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/services/dto"
)

type WorkspaceRepository struct {
	db *gorm.DB
}

func NewWorkspaceRepo(db *gorm.DB) *WorkspaceRepository {
	return &WorkspaceRepository{
		db: db,
	}
}

func (repo *WorkspaceRepository) Create(workspace *entities.Workspace, owner *entities.User) (*entities.Workspace, error) {
	dbUser := user.ToDBUser(owner)

	dbUser.ID = owner.ID

	logrus.Info(dbUser)

	dbWorkspace := ToDBWorkspace(workspace, dbUser)

	if err := repo.db.Create(dbWorkspace).Error; err != nil {
		return nil, err
	}

	return ToDomainWorkspace(dbWorkspace), nil
}

func (repo *WorkspaceRepository) FindAll(page, pageSize int, searchTerm string) ([]*entities.Workspace, error) {
	offset := (page - 1) * pageSize
	var workspaces []*models.Workspace

	query := repo.db.Offset(offset).Limit(pageSize)

	if searchTerm != "" {
		query = query.Where("name LIKE ?", "%"+searchTerm+"%")
	}

	if err := query.Find(&workspaces).Error; err != nil {
		return nil, err
	}

	return ToDomainWorkspaces(workspaces), nil
}

func (repo *WorkspaceRepository) FindById(id string) (*entities.Workspace, error) {
	var workspace *models.Workspace

	if err := repo.db.Model(&models.Workspace{}).Where("id = ?", id).First(&workspace).Error; err != nil {
		return nil, err
	}

	return ToDomainWorkspace(workspace), nil
}

func (repo *WorkspaceRepository) Update(workspaceID string, workspace *dto.WorkspaceDTO) error {
	result := repo.db.Model(&models.Workspace{}).Where("id = ?", workspaceID).Updates(workspace)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *WorkspaceRepository) Delete(workspaceID string) error {
	result := repo.db.Model(&models.Workspace{}).Delete(&models.Workspace{}, workspaceID)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
