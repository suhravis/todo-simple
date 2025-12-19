package models

import (
	"time"
	"todo-backend/database"
)

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
}

func GetAllTodos(db *database.DB) ([]Todo, error) {
	rows, err := db.Query("SELECT id, title, completed, created_at FROM todos ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, rows.Err()
}

func CreateTodo(db *database.DB, todo *Todo) error {
	err := db.QueryRow(
		"INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id, created_at",
		todo.Title, todo.Completed,
	).Scan(&todo.ID, &todo.CreatedAt)
	return err
}

func UpdateTodo(db *database.DB, todo *Todo) error {
	_, err := db.Exec(
		"UPDATE todos SET title = $1, completed = $2 WHERE id = $3",
		todo.Title, todo.Completed, todo.ID,
	)
	return err
}

func DeleteTodo(db *database.DB, id int) error {
	_, err := db.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}

