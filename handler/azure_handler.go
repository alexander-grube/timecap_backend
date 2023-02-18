package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/spctr-cc/backend-bugtrack/model"
	"github.com/spctr-cc/backend-bugtrack/utils"
)

func (h *Handler) CreateAzureTicket(c *fiber.Ctx) error {

	req := azureCreateTicketRequest{}
	t := &model.Ticket{}

	if err := req.bind(c, t, h.validator, h.accountStore); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	// if err := h.ticketStore.Create(t); err != nil {
	// 	return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	// }

	return c.Status(http.StatusCreated).JSON(newTicketCreatedResponse(t))
}
