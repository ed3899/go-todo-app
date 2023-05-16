package todo

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	PENDING  = "pending"
	PROGRESS = "in_progress"
	DONE     = "done"
)

type Todo struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type TodoRepository struct {
	database *sql.DB
}

func NewTodoRepository(database *sql.DB) *TodoRepository {
	return &TodoRepository{
		database,
	}
}

func (repository *TodoRepository) Create(todo Todo) (Todo, error) {
	_, err := repository.database.Exec("INSERT INTO todos (Name, Status) VALUES(?, ?);", todo.Name, todo.Status)
	if err != nil {
		err := fmt.Errorf("there was an error while inserting the todo '%#v'. The cause was: %v", todo.Name, err)
		return Todo{todo.Name, todo.Status}, err
	}

	log.Printf("%#v succesfully inserted", todo)
	return Todo{todo.Name, todo.Status}, nil
}
