package services

import "errors"

var (
	CannotCreateWorkspaceError = errors.New("cannot create workspace")
	ExistWorkspaceError        = errors.New("tag already exists")
	NotFoundWorkspaceError     = errors.New("tag not found")
	CannotUpdateWorkspaceError = errors.New("cannot update task")
	CannotDeleteWorkspaceError = errors.New("cannot delete task")
)

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
