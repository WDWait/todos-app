package service

import (
	"fmt"
	"todos-app/internal/model"
	"todos-app/internal/repository"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func NewTodoService() *TodoService {
	return &TodoService{
		repo: repository.NewTodoRepository(),
	}
}

func (s *TodoService) DeleteTodo(id int) error {
	return s.repo.Delete(id)
}

func (s *TodoService) UpdateTodo(id int, todo *model.Todo) error {
	if len(todo.Title) == 0 {
		return fmt.Errorf("title cannot be empty")
	}
	return s.repo.Update(id, todo)
}

func (s *TodoService) CreateTodo(todo *model.Todo) error {
	// 校验title 不能是空字符
	if len(todo.Title) == 0 {
		return fmt.Errorf("title cannot be empty")
	}
	return s.repo.Create(todo)
}

func (s *TodoService) GetTodoByID(id int) (*model.Todo, error) {
	return s.repo.GetByID(id)
}

func (s *TodoService) GetAllTodos() ([]model.Todo, error) {
	return s.repo.GetAll()
}
