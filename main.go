package main

import (
	"os"
	"time"

	models "github.com/spctr-cc/backend-bugtrack/Models"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var (
	PORT string = os.Getenv("PORT")
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
