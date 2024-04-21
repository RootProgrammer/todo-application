package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	todos := []ToDo{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("Everything is Fine!")
	})

	app.Post("/api/todos", handlePostRequest(&todos))

	app.Patch("/api/todos/:id/toggle", handleToggleRequest(todos))
    app.Patch("/api/todos/:id/done", handleMarkDoneRequest(todos))

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":4000"))
}
