package repository

import "todoApp/Entity"

type Authorization interface {
	CreateUser(user Entity.User) (int, error)
	GetUser(login Entity.Login) (Entity.User, error)
}

type TodoList interface {
	Create(userId int, list Entity.TodoList) (int, error)
}

type TodoItem interface {
}
