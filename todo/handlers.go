package todo

import (
	"database/sql"
	"log"

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
			"error": err,
		})
	}

	return c.Status(201).JSON(item)
}

func NewTodoHandler(repository *TodoRepository) *TodoHandler {
	return &TodoHandler{
		repository,
	}
}

func Register(router fiber.Router, database *sql.DB)  {
	todoRepository := NewTodoRepository(database)
	todoHandler := NewTodoHandler(todoRepository)

	todoRouter := router.Group("/todo")
	todoRouter.Post("/", todoHandler.Create)
}
