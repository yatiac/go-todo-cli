package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yatiac/go-todo-cli/models"
	"github.com/yatiac/go-todo-cli/repositories"
	"github.com/yatiac/go-todo-cli/services"
)

var todoRepo = repositories.NewJsonTodoRepository("todos.json")
var todoService = services.NewTodoService(*todoRepo)

func main() {
	// Display menu of options to the user
	println("Welcome to Go Todo CLI!")
	for {
		printOptions()
		println("Please enter your choice (1-6):")
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			println("Invalid input. Please enter a number between 1 and 6.")
			clearScreen()
			continue
		}

		switch choice {
		case 1:
			viewTodos()
		case 2:
			addTodo()
		case 3:
			toggleTodo()
		case 4:
			deleteTodo()
		case 5:
			addDescription()
		case 6:
			println("Exiting the application. Goodbye!")
			return
		default:
			println("Invalid choice. Please enter a number between 1 and 5.")
		}
	}
}

func viewTodos() {
	todos, err := todoService.GetAllTodos()
	if err != nil {
		println("Error fetching todos:", err.Error())
		return
	}
	if len(*todos) == 0 {
		println("No todos found.")
		return
	}
	println("Your Todos:")
	for i, todo := range *todos {
		status := " "
		if todo.Completed {
			status = "x"
		}
		fmt.Printf("%d. [%s] %s\n", i, status, todo.Title)
	}
}
func addTodo() {
	println("Enter todo title:")
	var title string
	reader := bufio.NewReader(os.Stdin)
	title, _ = reader.ReadString('\n')
	newTodo := &models.Todo{
		Title:       title,
		Description: "",
		Completed:   false,
	}
	err := todoService.CreateTodo(newTodo)
	if err != nil {
		println("Error adding todo:", err.Error())
		return
	}
	println("Todo added successfully!")
}
func toggleTodo() {
	println("Enter todo ID to toggle completion:")
	var id int
	_, err := fmt.Scan(&id)
	if err != nil {
		println("Invalid input. Please enter a valid todo ID.")
		return
	}
	err = todoService.ToggleTodo(id)
	if err != nil {
		println("Error toggling todo:", err.Error())
		return
	}
	println("Todo toggled successfully!")
}
func deleteTodo() {
	println("Enter todo ID to delete:")
	var id int
	_, err := fmt.Scan(&id)
	if err != nil {
		println("Invalid input. Please enter a valid todo ID.")
		return
	}
	err = todoService.DeleteTodo(id)
	if err != nil {
		println("Error deleting todo:", err.Error())
		return
	}
	println("Todo deleted successfully!")
}

func addDescription() {
	println("Enter todo ID to add description:")
	var id int
	_, err := fmt.Scan(&id)
	if err != nil {
		println("Invalid input. Please enter a valid todo ID.")
		return
	}
	println("Enter description:")
	var description string
	reader := bufio.NewReader(os.Stdin)
	description, _ = reader.ReadString('\n')

	err = todoService.AddDescription(id, description)
	if err != nil {
		println("Error adding description:", err.Error())
		return
	}
	println("Description added successfully!")
}

func printOptions() {
	println("1. View Todos")
	println("2. Add Todo")
	println("3. Toggle Todo Completion")
	println("4. Delete Todo")
	println("5. Add Description to Todo")
	println("6. Exit")
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
