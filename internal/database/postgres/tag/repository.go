package tag

import (
	"gorm.io/gorm"
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/services/dto"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepo(db *gorm.DB) *TagRepository {
	return &TagRepository{
		db: db,
	}
}

func (t TagRepository) GetById(id string) (*entities.Tag, error) {
	var tag *Tag

	err := t.db.Model(&tag).Where("id = ?", id).Find(&tag).Error

	if err != nil {
		return nil, err
	}

	return ToDomainTag(tag), nil
}

func (t TagRepository) GetByName(name string) (*entities.Tag, error) {
	var tag *Tag

	result := t.db.Model(&Tag{}).Where("name = ?", name).Find(&tag)

	if result.Error != nil {
		return nil, result.Error
	}

	return ToDomainTag(tag), nil
}

func (t TagRepository) GetAll(searchTerm string) ([]*entities.Tag, error) {
	var tags []*Tag

	query := t.db.Model(&Tag{})

	if searchTerm != "" {
		query = query.Where("name LIKE ?", "%"+searchTerm+"%")
	}

	result := query.Find(&tags)

	if result.Error != nil {
		return nil, result.Error
	}

	return ToDomainTags(tags), nil
}

func (t TagRepository) Create(tag *entities.Tag) (*entities.Tag, error) {
	newTag := ToDBTag(tag)

	if err := t.db.Model(&Tag{}).Create(newTag).Error; err != nil {
		return nil, err
	}

	return ToDomainTag(newTag), nil
}

func (t TagRepository) Update(id string, tag *dto.Tag) error {
	result := t.db.Model(&Tag{}).Where("id = ?", id).Updates(tag)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t TagRepository) Delete(id string) error {
	result := t.db.Model(&Tag{}).Where("id = ?", id).Delete(&Tag{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
