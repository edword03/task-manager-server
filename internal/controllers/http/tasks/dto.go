package tasks

type TaskDTO struct {
	Title    string   `validate:"required,alphanum" json:"title"`
	Content  string   `validate:"omitempty,alphanum" json:"content"`
	FromTime string   `validate:"omitempty,date" json:"fromTime"`
	ToTime   string   `validate:"omitempty,date" json:"toTime"`
	Priority string   `validate:"omitempty,num" json:"priority"`
	Author   string   `validate:"required,alphanum" json:"author"`
	Assignee string   `validate:"omitempty,alphanum" json:"assignee"`
	Tags     []string `validate:"omitempty,dive,tag" json:"tags"`
}
