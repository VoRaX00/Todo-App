package repository

import "todoApp/Entity"

type Authorization interface {
	Create(user Entity.User) (int, error)
	Get(login Entity.Login) (Entity.User, error)
}

type TodoList interface {
	Create(userId int, list Entity.TodoList) (int, error)
	GetAll(userId int) ([]Entity.TodoList, error)
	GetById(userId, listId int) (Entity.TodoList, error)
	Update(userId, listId int, list Entity.UpdateList) error
	Delete(userId, listId int) error
}

type TodoItem interface {
	Create(listId int, item Entity.TodoItem) (int, error)
	GetAll(userId, listId int) ([]Entity.TodoItem, error)
	GetById(userId, itemId int) (Entity.TodoItem, error)
	Update(userId, itemId int, item Entity.UpdateItem) error
	Delete(userId, itemId int) error
}
