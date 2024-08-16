package repository

import "todoApp/Entity"

type Authorization interface {
	CreateUser(user Entity.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}
