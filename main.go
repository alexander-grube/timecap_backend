package main

import (
	models "spctr/bugtrack/models"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const (
	PORT string = ":8000"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		ticket := models.Ticket{
			Topic:     "Test",
			Timestamp: time.Now().UnixMilli(),
		}
		return c.JSON(ticket)
	})

	app.Listen(PORT)
}
