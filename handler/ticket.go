package handler

import (
	"net/http"
	"strconv"

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

func (h *Handler) GetTicketByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	t, err := h.ticketStore.GetByID(uint(idUint))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(utils.NewError(err))
	}

	return c.JSON(newTicketResponse(t))
}

func (h *Handler) GetAllTickets(c *fiber.Ctx) error {
	t, err := h.ticketStore.GetAll()
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(utils.NewError(err))
	}

	return c.JSON(t)
}
