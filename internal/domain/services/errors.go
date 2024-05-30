package services

import "errors"

var (
	ExistTagError    = errors.New("tag already exists")
	NotFoundTagError = errors.New("tag not found")
)

var (
	ExistsTaskError       = errors.New("task already exists")
	NotFoundTaskError     = errors.New("task not found")
	CannotCreateTaskError = errors.New("cannot create task")
	CannotUpdateTaskError = errors.New("cannot update task")
	CannotDeleteTaskError = errors.New("cannot delete task")
)
