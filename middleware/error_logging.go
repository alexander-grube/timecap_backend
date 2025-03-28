package middleware

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func LogRequestBodyOnError(c fiber.Ctx) error {
	// Read and store the request body
	body := c.Body()

	// Process the request
	err := c.Next()

	// Check for errors or 4xx and 5xx status codes
	status := c.Response().StatusCode()
	if err != nil || status >= fiber.StatusBadRequest {
		log.Printf("Error: %v, Status: %d, Request Body: %s",
			err, status, string(body))
	}

	return err
}