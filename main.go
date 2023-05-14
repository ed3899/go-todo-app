package main

import (
	"log"
	"github.com/edca3899/go-todo-mysql/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	db.ConnectDB()
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("App running")
	})

	log.Fatal(app.Listen(":8080"))
}