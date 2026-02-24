package service

import (
	"newtodo"
	"newtodo/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list newtodo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
