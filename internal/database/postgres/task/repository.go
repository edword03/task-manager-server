package task

import (
	"gorm.io/gorm"
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/services/dto"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) TaskRepository {
	return TaskRepository{
		db: db,
	}
}

func (t TaskRepository) Create(task *entities.Task) (*Task, error) {
	dbTask := ToDBTask(task)
	if err := t.db.Model(&Task{}).Create(&dbTask).Error; err != nil {
		return nil, err
	}

	return dbTask, nil
}

func (t TaskRepository) GetById(id string) (*entities.Task, error) {
	var task *Task

	if err := t.db.Model(&Task{}).Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}

	return ToDomainTask(task), nil
}

func (t TaskRepository) GetAll(page, pageSize int, searchTerm string) ([]*entities.Task, error) {
	offset := (page - 1) * pageSize
	var tasks []*Task

	query := t.db.Model(&Task{}).Offset(offset).Limit(pageSize)

	if searchTerm != "" {
		query = query.Where("name LIKE ?", "%"+searchTerm+"%")
	}

	if err := query.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return ToDomainTasks(tasks), nil
}

func (t TaskRepository) Update(taskId string, task *dto.TaskDTO) error {
	result := t.db.Model(&Task{}).Where("id = ?", taskId).Updates(task)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t TaskRepository) Delete(taskId string) error {
	result := t.db.Model(&Task{}).Delete(&Task{}, taskId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
