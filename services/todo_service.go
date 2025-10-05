package services

import (
	"github.com/yatiac/go-todo-cli/models"
	"github.com/yatiac/go-todo-cli/repositories"
)

type TodoService struct {
	repo repositories.JsonTodoRepository
}

func NewTodoService(repo repositories.JsonTodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) GetAllTodos() (*[]models.Todo, error) {
	return s.repo.GetAllTodos()
}
func (s *TodoService) CreateTodo(todo *models.Todo) error {
	return s.repo.CreateTodo(todo)
}
func (s *TodoService) ToggleTodo(id int) error {
	return s.repo.ToggleTodo(id)
}
func (s *TodoService) DeleteTodo(id int) error {
	return s.repo.DeleteTodo(id)
}
func (s *TodoService) AddDescription(id int, description string) error {
	return s.repo.AddDescription(id, description)
}
