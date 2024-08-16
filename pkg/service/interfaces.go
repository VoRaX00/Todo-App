package service

import "todoApp/Entity"

type Authorization interface {
	CreateUser(user Entity.User) (int, error)
	GenerateToken(login Entity.Login) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoList interface {
	Create(userId int, list Entity.TodoList) (int, error)
	GetAll(userId int) ([]Entity.TodoList, error)
	GetById(userId int, listId int) (Entity.TodoList, error)
	Update(userId int, listId int, list Entity.UpdateListInput) error
	Delete(userId int, listId int) error
}

type TodoItem interface {
}
