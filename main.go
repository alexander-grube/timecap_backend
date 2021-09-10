package main

import (
	"log"
	"os"
	"time"

	model "github.com/spctr-cc/backend-bugtrack/model"

	"database/sql"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var (
	PORT string = ":" + os.Getenv("PORT")
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		ticket := model.Ticket{
			Topic:     "Test",
			Timestamp: time.Now().UnixMilli(),
		}
		return c.JSON(ticket)
	})

	account := model.Account{
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@example.com",
		Password:  "johnIsADoe",
		Mobile:    "555-12345678",
		Street:    "John Doe Street",
		Zipcode:   "12345",
		City:      "John Doe City",
		Country:   "Canada",
	}

	app.Post("/account", func(c *fiber.Ctx) error {
		return c.JSON(account)
	})

	app.Get("/db", func(c *fiber.Ctx) error {
		sel := "SELECT * from accounts"
		row, err := db.Query(sel)
		if err != nil{
			log.Fatal(err)
		}
		col, err := row.Columns()
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(col)
	})

	app.Listen(PORT)
}
