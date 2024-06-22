package tag

import (
	"task-manager/internal/database/postgres/models"
	"task-manager/internal/domain/entities"
)

func ToDomainTag(tag *models.Tag) *entities.Tag {
	return &entities.Tag{
		ID:    tag.ID,
		Name:  tag.Name,
		Color: tag.Color,
	}
}

func ToDBTag(tag *entities.Tag) *models.Tag {
	return &models.Tag{
		ID:    tag.ID,
		Name:  tag.Name,
		Color: tag.Color,
	}
}

func ToDomainTags(tags []*models.Tag) []*entities.Tag {
	var result []*entities.Tag

	for _, tag := range tags {
		result = append(result, ToDomainTag(tag))
	}

	return result
}
