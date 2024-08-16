package repository

import "todoApp/Entity"

type Authorization interface {
	CreateUser(user Entity.User) (int, error)
	GetUser(login Entity.Login) (Entity.User, error)
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
