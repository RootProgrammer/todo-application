// mark_done_handler.go
package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
)

func handleMarkDoneRequest(todos []ToDo) fiber.Handler {
    return func(c *fiber.Ctx) error {
        id, err := c.ParamsInt("id")
        if err != nil {
            log.Println("Error extracting ID:", err)
            return c.Status(400).SendString("Invalid ID")
        }

        log.Println("Received PATCH request to mark todo as done with ID:", id)

        for index, todo := range todos {
            if todo.Id == id {
                todos[index].Done = true
                log.Println("Todo marked as done successfully")
                return c.JSON(fiber.Map{
                    "message": "Todo marked as done successfully",
                })
            }
        }

        log.Println("Todo with ID not found:", id)
        return c.Status(404).SendString("Todo not found")
    }
}