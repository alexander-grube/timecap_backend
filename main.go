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

	// account := model.Account{
	// 	Firstname: "John",
	// 	Lastname:  "Doe",
	// 	Email:     "john.doe@example.com",
	// 	Password:  "johnIsADoe",
	// 	Mobile:    "555-12345678",
	// 	Street:    "John Doe Street",
	// 	Zipcode:   "12345",
	// 	City:      "John Doe City",
	// 	Country:   "Canada",
	// }

	app.Post("/account", func(c *fiber.Ctx) error {
		ins := "INSERT INTO accounts VALUES('John', 'Doe', 'john.doe@example.com', 'johnIsADoe', '555-12345678', 'John Doe Street', '12345', 'John Doe City', 'Canada'"
		res, err := db.Exec(ins)
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(res)
	})

	app.Get("/db", func(c *fiber.Ctx) error {
		sel := "SELECT * from accounts"
		res, err := db.Exec(sel)
		if err != nil{
			log.Fatal(err)
		}
		return c.JSON(res)
	})

	app.Listen(PORT)
}
