package main

import "github.com/gofiber/fiber/v2"

func handlePostRequest(todos *[]ToDo) fiber.Handler {
	return func(c *fiber.Ctx) error {
		todo := &ToDo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		todo.Id = len(*todos) + 1

		*todos = append(*todos, *todo)

		return c.JSON(*todos)
	}
}
