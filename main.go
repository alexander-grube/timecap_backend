package main

import (
	"os"
	"time"

	models "github.com/spctr-cc/backend-bugtrack/Models"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var (
	PORT string = ":" + os.Getenv("PORT")
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

	account := models.Account{
		Firstname: "John",
		Lastname: "Doe",
		Email: "john.doe@example.com",
		Password: "johnIsADoe",
		Mobile: "555-12345678",
		Street: "John Doe Street",
		Zipcode: "12345",
		City: "John Doe City",
		Country: "Canada",
	}

	app.Post("/account", func(c *fiber.Ctx) error {
		return c.JSON(account)
	})

	app.Listen(PORT)
}
