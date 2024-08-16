package repository

import "todoApp/Entity"

type Authorization interface {
	CreateUser(user Entity.User) (int, error)
	GetUser(login Entity.Login) (Entity.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}
