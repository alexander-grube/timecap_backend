package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateAzureTicket(c *fiber.Ctx) error {
	log.Println("CreateAzureTicket")
	log.Println(string(c.Body()))
	return nil
}
