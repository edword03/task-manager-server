package services

import "errors"

var (
	ExistTagError    = errors.New("tag already exists")
	NotFoundTagError = errors.New("tag not found")
)
