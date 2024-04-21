// toggle_handler.go
package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
)

func handleToggleRequest(todos []ToDo) fiber.Handler {
    return func(c *fiber.Ctx) error {
        id, err := c.ParamsInt("id")
        if err != nil {
            log.Println("Error extracting ID:", err)
            return c.Status(400).SendString("Invalid ID")
        }

        log.Println("Received PATCH request to toggle todo with ID:", id)

        for index, todo := range todos {
            if todo.Id == id {
                todos[index].Done = !todos[index].Done
                log.Println("Todo status toggled successfully")
                return c.JSON(fiber.Map{
                    "message": "Todo status toggled successfully",
                })
            }
        }

        log.Println("Todo with ID not found:", id)
        return c.Status(404).SendString("Todo not found")
    }
}