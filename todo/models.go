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

type TodoResponse struct {
	TodoId interface{} `json:",omitempty"`
	Todo
}

type TodoRepository struct {
	database *sql.DB
}

func NewTodoRepository(database *sql.DB) *TodoRepository {
	return &TodoRepository{
		database,
	}
}

func (repository *TodoRepository) Create(todo Todo) (TodoResponse, error) {
	result, err := repository.database.Exec("INSERT INTO todos (Name, Status) VALUES(?, ?);", todo.Name, todo.Status)
	if err != nil {
		err := fmt.Errorf("there was an error while inserting the todo '%#v'. The cause was: %v", todo, err)
		return TodoResponse{nil, todo}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		err := fmt.Errorf("there was an error while retrieving the inserted id of todo: %#v. The cause was: %v", todo, err)
		return TodoResponse{nil, todo}, err
	}

	log.Printf("%#v succesfully inserted", todo)
	return TodoResponse{id, todo}, nil
}

func (repository *TodoRepository) Find(id int) (TodoResponse, error) {
	var todo TodoResponse
	sqlQuery := `
	SELECT * FROM todos
	WHERE TodoId = ?;
	`

	switch err := repository.database.QueryRow(sqlQuery, id).Scan(&todo.TodoId, &todo.Name, &todo.Status); {
	case err == sql.ErrNoRows:
		err := fmt.Errorf("the todo with id: %d was not found", id)
		log.Print(err)

		return TodoResponse{}, err
	case err != nil:
		err := fmt.Errorf("something bad happened while looking for Todo with id:%d which has structure %#v. Cause: %v", id, todo, err)
		log.Print(err)

		return TodoResponse{}, err
	default:
		log.Printf("Sending %#v to requester", todo)

		return todo, nil
	}
}
