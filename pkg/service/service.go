package service

import (
	"todoApp/pkg/repository"
)

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(rep *repository.Repository) *Service {
	return &Service{}
}
