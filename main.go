package main

import (
	model "spctr/bugtrack/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	PORT string = ":8000"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		ticket := model.Ticket{
			Topic:     "Test",
			Timestamp: time.Now().UnixMilli(),
		}
		return c.JSON(ticket)
	})

	app.Listen(PORT)
}
