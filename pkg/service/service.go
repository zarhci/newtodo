package service

import (
	"newtodo"
	"newtodo/pkg/repository"
)

type Authorization interface {
	CreateUser(user newtodo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list newtodo.TodoList) (int, error)
	GetAll(userId int) ([]newtodo.TodoList, error)
	GetById(userId, listId int) (newtodo.TodoList, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
