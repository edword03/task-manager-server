package tag

import "task-manager/internal/domain/entities"

func ToDomainTag(tag *Tag) *entities.Tag {
	return &entities.Tag{
		ID:    tag.ID,
		Name:  tag.Name,
		Color: tag.Color,
	}
}

func ToDBTag(tag *entities.Tag) *Tag {
	return &Tag{
		ID:    tag.ID,
		Name:  tag.Name,
		Color: tag.Color,
	}
}

func ToDomainTags(tags []*Tag) []*entities.Tag {
	var result []*entities.Tag

	for _, tag := range tags {
		result = append(result, ToDomainTag(tag))
	}

	return result
}
