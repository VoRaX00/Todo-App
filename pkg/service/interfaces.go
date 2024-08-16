package service

import "todoApp/Entity"

type Authorization interface {
	CreateUser(user Entity.User) (int, error)
	GenerateToken(login Entity.Login) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoList interface {
	Create(userId int, list Entity.TodoList) (int, error)
}

type TodoItem interface {
}
