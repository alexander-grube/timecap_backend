package main

import (
	"log"
	"net/http"
	"os"
	"time"

	model "github.com/spctr-cc/backend-bugtrack/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

var (
	PORT string = ":" + os.Getenv("PORT")
	db   *gorm.DB
	err  error
)

func main() {
	db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Connected")
	db.AutoMigrate(&model.Account{})
	db.AutoMigrate(&model.Ticket{})
	log.Println("Database Migrated")

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
		Email:     "john.doe@example4.com",
		Password:  "johnIsADoe",
		Mobile:    "555-12345678",
		Street:    "John Doe Street",
		Zipcode:   "12345",
		City:      "John Doe City",
		Country:   "Canada",
	}

	app.Post("/account", func(c *fiber.Ctx) error {
		createdAccount := db.Create(&account)
		err = createdAccount.Error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err)
		}
		return c.JSON(&account)
	})

	app.Get("/account/:id", func(c *fiber.Ctx) error {
		var account model.Account
		db.First(&account, c.Params("id"))
		if account.Email == "" {
			return c.Status(http.StatusNotFound).SendString("Account not found")
		}
		return c.JSON(account)
	})

	app.Listen(PORT)
}
