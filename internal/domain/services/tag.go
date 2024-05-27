package services

import (
	"errors"
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/services/dto"
)

type tagRepository interface {
	GetById(id string) (*entities.Tag, error)
	GetByName(name string) (*entities.Tag, error)
	GetAll(searchTerm string) ([]*entities.Tag, error)
	Create(tag *entities.Tag) (*entities.Tag, error)
	Update(id string, tag *dto.Tag) error
	Delete(id string) error
}

type TagService struct {
	tagRepository tagRepository
}

func NewTagService(repo tagRepository) *TagService {
	return &TagService{
		tagRepository: repo,
	}
}

func (t TagService) CreateTag(tag *dto.Tag) (*entities.Tag, error) {
	existTag, err := t.tagRepository.GetByName(tag.Name)

	if err != nil {
		return nil, err
	}

	if existTag != nil && existTag.ID != "" {
		return nil, ExistTagError
	}

	var color = "#000"

	if tag.Color != "" {
		color = tag.Color
	}

	newTag := &entities.Tag{
		Name:  tag.Name,
		Color: color,
	}

	newTag, err = t.tagRepository.Create(newTag)

	if err != nil {
		return nil, err
	}

	return newTag, err
}

func (t TagService) GetTagById(id string) (*entities.Tag, error) {
	tag, err := t.tagRepository.GetById(id)

	if err != nil {
		return nil, err
	}

	if tag.ID == "" {
		return nil, errors.New("tag not found")
	}

	return tag, nil
}

func (t TagService) GetTagByName(name string) (*entities.Tag, error) {
	tag, err := t.tagRepository.GetByName(name)
	if err != nil {
		return nil, err
	}

	if tag.ID == "" {
		return nil, errors.New("tag not found")
	}

	return tag, nil
}

func (t TagService) GetTags(searchTerm string) ([]*entities.Tag, error) {
	tags, err := t.tagRepository.GetAll(searchTerm)

	if err != nil {
		return nil, err
	}

	if len(tags) == 0 {
		return nil, NotFoundTagError
	}

	return tags, nil
}

func (t TagService) UpdateTag(id string, tag *dto.Tag) error {
	currentTag, err := t.tagRepository.GetById(id)

	if err != nil {
		return err
	}

	if currentTag == nil {
		return NotFoundTagError
	}

	if err := t.tagRepository.Update(id, tag); err != nil {
		return err
	}

	return nil
}

func (t TagService) DeleteTagById(id string) error {
	tag, err := t.tagRepository.GetById(id)

	if err != nil {
		return err
	}

	if tag.ID == "" {
		return NotFoundTagError
	}

	if err := t.tagRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
