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

func (repository *TodoRepository) Find(id int) (Todo, error) {
	var todo Todo
	var extra string
	sqlQuery := `
	SELECT * FROM todos
	WHERE TodoId = ?;
	`

	switch err := repository.database.QueryRow(sqlQuery, id).Scan(&extra ,&todo.Name, &todo.Status); {
	case err == sql.ErrNoRows:
		err := fmt.Errorf("the todo with id: %d was not found", id)
		log.Print(err)

		return Todo{}, err
	case err != nil:
		err := fmt.Errorf("something bad happened while looking for Todo with id:%d which has structure %#v. Cause: %v", id, todo, err)
		log.Print(err)

		return Todo{}, err
	default:
		log.Printf("Sending %#v to requester", todo)

		return todo, nil
	}
}
