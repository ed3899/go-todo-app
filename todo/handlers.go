package todo

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	repository *TodoRepository
}

func (handler *TodoHandler) Create(c *fiber.Ctx) error {
	data := new(Todo)

	if err := c.BodyParser(data); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "error": err})
	}

	log.Printf("Getting body %#v", *data)

	item, err := handler.repository.Create(*data)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"message": "Failed creating item",
			"error":   err,
		})
	}

	return c.Status(201).JSON(item)
}

func (handler *TodoHandler) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		err := fmt.Errorf("there was an error parsing the parameter id: %v. Cause is: %#v", id, err)
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	todo, err := handler.repository.Find(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(todo)
}

func (handler *TodoHandler) GetAll(c *fiber.Ctx) error {
	todos, err := handler.repository.FindAll()
	if err != nil {
		err := fmt.Errorf("there was an error while finding all todos: %#v", err)
		return c.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}

	log.Printf("from get all handler %#v", todos)

	return c.Status(200).JSON(todos)
}

func (handler *TodoHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		err := fmt.Errorf("there was an error while parsing your request. Cause: %#v", err)
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	todo, err := handler.repository.Find(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	todoData := new(Todo)

	if err := c.BodyParser(todoData); err != nil {
		err := fmt.Errorf("there was an error while parsing your request. Cause: %#v", err)
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	todo.Name = todoData.Name
	todo.Status = todoData.Status

	item, err := handler.repository.Save(todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(item)
}

func NewTodoHandler(repository *TodoRepository) *TodoHandler {
	return &TodoHandler{
		repository,
	}
}

func Register(router fiber.Router, database *sql.DB) {
	todoRepository := NewTodoRepository(database)
	todoHandler := NewTodoHandler(todoRepository)

	todoRouter := router.Group("/todo")
	todoRouter.Get("/:id", todoHandler.Get)
	todoRouter.Get("/", todoHandler.GetAll)
	todoRouter.Post("/", todoHandler.Create)
	todoRouter.Put("/:id", todoHandler.Update)
}
