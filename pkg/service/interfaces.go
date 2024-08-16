package service

import "todoApp/Entity"

type Authorization interface {
	CreateUser(user Entity.User) (int, error)
	GenerateToken(login Entity.Login) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}
