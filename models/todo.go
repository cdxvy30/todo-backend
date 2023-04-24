package models

import (
	"time"
	db "todo-backend/utils"
)

type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetAllTodos() ([]Todo, error) {
	db, err := db.ConnectToDB()

	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}

	todos := []Todo{}
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func GetTodoById(id string) (Todo, error) {
	db, err := db.ConnectToDB()

	var todo Todo
	err = db.QueryRow("SELECT id, title, description, completed FROM todos WHERE id = $1", id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func CreateTodo(todo *Todo) error {
	db, err := db.ConnectToDB()

	_, err = db.Exec("INSERT INTO todos(title, description, completed) VALUES($1, $2, $3)", todo.Title, todo.Description, todo.Completed)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTodoById(id string, todo *Todo) error {
	db, err := db.ConnectToDB()

	_, err = db.Exec("UPDATE todos SET title = $1, description = $2, completed = $3 WHERE id = $4", todo.Title, todo.Description, todo.Completed, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTodoById(id string) error {
	db, err := db.ConnectToDB()

	_, err = db.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func atoi(s string) int {
	i := 0
	for _, c := range s {
		i = i*10 + int(c-'0')
	}
	return i
}
