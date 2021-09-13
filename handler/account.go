package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/spctr-cc/backend-bugtrack/model"
	"github.com/spctr-cc/backend-bugtrack/utils"
)

func (h *Handler) SignUp(c *fiber.Ctx) error {
	var a model.Account
	req := accountRegisterRequest{}

	if err := req.bind(c, &a, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	if err := h.accountStore.Create(&a); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusCreated).JSON(NewAccountResponse(&a))
}

func (h *Handler) CurrentAccount(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}
	a, err := h.accountStore.GetByID(uint(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if a == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound())
	}
	return c.Status(http.StatusOK).JSON(NewAccountResponse(a))
}
