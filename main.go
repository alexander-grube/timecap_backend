package main

import (
	"fmt"
	"os"

	"github.com/spctr-cc/backend-bugtrack/db"
	"github.com/spctr-cc/backend-bugtrack/handler"
	"github.com/spctr-cc/backend-bugtrack/store"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	PORT string = ":" + os.Getenv("PORT")
	err  error
)

func main() {
	
	d := db.New()
	fmt.Println("Database Connected")
	db.AutoMigrate(d)
	fmt.Println("Database Migrated")

	app := fiber.New()

	app.Use(logger.New())

	app.Use(cors.New())

	as := store.NewAccountStore(d)

	ts := store.NewTicketStore(d)

	h := handler.NewHandler(as, ts)

	h.Register(app)

	app.Listen(PORT)
}
