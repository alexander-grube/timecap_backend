package main

import (
	"encoding/json"
	"fmt"
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
	fmt.Println("Database Connected")
	db.AutoMigrate(&model.Account{})
	db.AutoMigrate(&model.Ticket{})
	fmt.Println("Database Migrated")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		ticket := model.Ticket{
			Topic:     "Test",
			Timestamp: time.Now().UnixMilli(),
		}
		return c.JSON(ticket)
	})

	app.Post("/account", func(c *fiber.Ctx) error {
		var account model.Account
	    err = json.Unmarshal(c.Body(), &account)
		if err != nil {
			log.Fatal("Could not parse account JSON", err)
		}
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
			return c.Status(http.StatusNotFound).SendString("Account not found. Could not get.")
		}
		return c.JSON(account)
	})

	app.Post("/account/delete/:id", func (c *fiber.Ctx) error {
		var account model.Account
		db.First(&account, c.Params("id"))
		db.Delete(&account)
		if account.Email == "" {
			return c.Status(http.StatusNotFound).SendString("Account not found. Could not delete.")
		}
		return c.JSON(&account)
	})

	app.Listen(PORT)
}
