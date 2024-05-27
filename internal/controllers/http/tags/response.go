package tags

import "task-manager/internal/domain/entities"

type TagResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

func toTagResp(tag *entities.Tag) *TagResponse {
	return &TagResponse{
		ID:    tag.ID,
		Name:  tag.Name,
		Color: tag.Color,
	}
}
