package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/spctr-cc/backend-bugtrack/model"
	"github.com/spctr-cc/backend-bugtrack/utils"
)

func (h *Handler) CreateTicket(c *fiber.Ctx) error {
	var t model.Ticket
	req := ticketCreateRequest{}

	if err := req.bind(c, &t, h.validator, h.accountStore); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	if err := h.ticketStore.Create(&t); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusCreated).JSON(newTicketResponse(&t))
}
