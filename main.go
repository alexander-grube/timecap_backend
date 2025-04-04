package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bytedance/sonic"
	"github.com/joho/godotenv"
	"github.com/spctr-cc/backend-bugtrack/db"
	"github.com/spctr-cc/backend-bugtrack/handler"
	"github.com/spctr-cc/backend-bugtrack/middleware"
	"github.com/spctr-cc/backend-bugtrack/store"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
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

	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	app.Use(logger.New())

	app.Use(cors.New())

	// Use the middleware before other handlers
	app.Use(middleware.LogRequestBodyOnError)

	as := store.NewAccountStore(d)

	ts := store.NewTicketStore(d)

	ps := store.NewProjectStore(d)

	azs := store.NewAzureStore()

	h := handler.NewHandler(as, ts, ps, azs)

	h.Register(app)

	app.Listen(PORT)
}
