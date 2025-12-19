package models

import (
	"errors"
	"sync"
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
}

var (
	todos  []Todo
	nextID int
	mu     sync.RWMutex
)

func init() {
	todos = make([]Todo, 0)
	nextID = 1
}

func GetAllTodos() ([]Todo, error) {
	mu.RLock()
	defer mu.RUnlock()
	
	result := make([]Todo, len(todos))
	copy(result, todos)
	return result, nil
}

func CreateTodo(todo *Todo) error {
	mu.Lock()
	defer mu.Unlock()
	
	todo.ID = nextID
	nextID++
	todo.CreatedAt = time.Now()
	todos = append(todos, *todo)
	return nil
}

func UpdateTodo(todo *Todo) error {
	mu.Lock()
	defer mu.Unlock()
	
	for i := range todos {
		if todos[i].ID == todo.ID {
			todos[i].Title = todo.Title
			todos[i].Completed = todo.Completed
			*todo = todos[i]
			return nil
		}
	}
	return errors.New("todo not found")
}

func DeleteTodo(id int) error {
	mu.Lock()
	defer mu.Unlock()
	
	for i := range todos {
		if todos[i].ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}
