package service

import (
	"todoApp/pkg/handler"
	"todoApp/pkg/repository"
)

type Service struct {
	handler.Authorization
	handler.TodoList
	handler.TodoItem
}

func NewService(rep *repository.Repository) *Service {
	return &Service{}
}
