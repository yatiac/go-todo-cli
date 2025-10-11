package repositories

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/yatiac/go-todo-cli/models"
)

type TodoRepository interface {
	GetAllTodos() (*[]models.Todo, error)
	CreateTodo(todo *models.Todo) error
	ToggleTodo(index int) error
	DeleteTodo(id int) error
}

// from json file
type JsonTodoRepository struct {
	filePath string
	todos    *[]models.Todo
}

func NewJsonTodoRepository(filePath string) *JsonTodoRepository {
	return &JsonTodoRepository{filePath: filePath, todos: &[]models.Todo{}}
}

func (r *JsonTodoRepository) GetAllTodos() (*[]models.Todo, error) {
	if _, error := os.Stat(r.filePath); error == nil {
		data, err := os.ReadFile(r.filePath)
		if err != nil {
			log.Fatal("Error reading the todos file")
		}
		json.Unmarshal(data, &r.todos)
	}
	return r.todos, nil
}

func (r *JsonTodoRepository) CreateTodo(todo *models.Todo) error {
	*r.todos = append(*r.todos, *todo)
	return r.saveTodos()
}

func (r *JsonTodoRepository) ChangeStatus(status models.TodoStatus, id int) error {
	if id < 0 || id >= len(*r.todos) {
		fmt.Println("Todo not found")
		return nil
	}
	(*r.todos)[id].Status = status.String()
	return r.saveTodos()
}

func (r *JsonTodoRepository) DeleteTodo(id int) error {
	if id < 0 || id >= len(*r.todos) {
		fmt.Println("Todo not found")
		return nil
	}
	*r.todos = append((*r.todos)[:id], (*r.todos)[id+1:]...)
	return r.saveTodos()
}

func (r *JsonTodoRepository) AddDescription(id int, description string) error {
	if id < 0 || id >= len(*r.todos) {
		fmt.Println("Todo not found")
		return nil
	}
	(*r.todos)[id].Description = description
	return r.saveTodos()
}

func (r *JsonTodoRepository) saveTodos() error {
	data, err := json.MarshalIndent(*r.todos, "", "  ")
	if err != nil {
		log.Fatal("Error marshalling todos to JSON")
	}
	err = os.WriteFile(r.filePath, data, 0644)
	if err != nil {
		log.Fatal("Error writing todos to file")
	}
	return nil
}
