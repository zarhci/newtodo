package repository

import (
	"newtodo"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user newtodo.User) (int, error)
	GetUser(username, password string) (newtodo.User, error)
}

type TodoList interface {
	Create(userId int, list newtodo.TodoList) (int, error)
	GetAll(userId int) ([]newtodo.TodoList, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
