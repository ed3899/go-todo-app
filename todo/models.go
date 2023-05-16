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
	TodoId int64 `json:",omitempty"`
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
		return TodoResponse{0, todo}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		err := fmt.Errorf("there was an error while retrieving the inserted id of todo: %#v. The cause was: %v", todo, err)
		return TodoResponse{0, todo}, err
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

func (repository *TodoRepository) FindAll() ([]TodoResponse, error) {
	// get rows from db
	sqlQuery := `
	SELECT * FROM todos;
	`
	rows, err := repository.database.Query(sqlQuery)
	if err != nil {
		return []TodoResponse{}, nil
	}
	defer rows.Close()

	// iterate through rows and fill todo array
	var todoResponses []TodoResponse
	for rows.Next() {
		var todoResponse TodoResponse
		if err := rows.Scan(&todoResponse.TodoId, &todoResponse.Name, &todoResponse.Status); err != nil {
			err := fmt.Errorf("there was an error scanning the row. Cause: %#v", err)
			log.Print(err)
			return todoResponses, err
		}
		todoResponses = append(todoResponses, todoResponse)
	}

	if err = rows.Err(); err != nil {
		err := fmt.Errorf("an error was encountered during the iteration: %#v", err)
		return todoResponses, err
	}

	// return todo array if no errors before
	return todoResponses, nil
}
