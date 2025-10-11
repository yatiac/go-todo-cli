package models

import "time"

type Todo struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
}

type TodoStatus int

const (
	Pending TodoStatus = iota
	InProgress
	Completed
)

func (s TodoStatus) String() string {
	return [...]string{"Pending", "In Progress", "Completed"}[s]
}
