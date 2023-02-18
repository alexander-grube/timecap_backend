package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spctr-cc/backend-bugtrack/db"
	"github.com/spctr-cc/backend-bugtrack/handler"
	"github.com/spctr-cc/backend-bugtrack/store"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-l" {
		// read from .env file
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	PORT := ":" + os.Getenv("PORT")

	d := db.New()
	fmt.Println("Database Connected")
	db.AutoMigrate(d)
	fmt.Println("Database Migrated")

	app := fiber.New()

	app.Use(logger.New())

	app.Use(cors.New())

	as := store.NewAccountStore(d)

	ts := store.NewTicketStore(d)

	ps := store.NewProjectStore(d)

	azs := store.NewAzureStore()

	h := handler.NewHandler(as, ts, ps, azs)

	h.Register(app)

	if len(os.Args) > 1 && os.Args[1] == "-l" {
		app.Listen(PORT)
	}

	app.ListenTLS(os.Getenv("IP")+PORT, "/etc/letsencrypt/live/backend-bugtrack.alexgrube.dev/fullchain.pem", "/etc/letsencrypt/live/backend-bugtrack.alexgrube.dev/privkey.pem")
}
