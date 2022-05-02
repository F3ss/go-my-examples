package service

import "todo/pkg/repositories"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{
		repo.Authorization,
		repo.TodoList,
		repo.TodoItem,
	}
}
