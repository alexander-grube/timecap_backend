package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/spctr-cc/backend-bugtrack/model"
	"github.com/spctr-cc/backend-bugtrack/utils"
)

func (h *Handler) CreateProject(c *fiber.Ctx) error {
	var p model.Project
	req := projectCreateRequest{}

	if err := req.bind(c, &p, h.validator, h.accountStore); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	if err := h.projectStore.Create(&p); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusCreated).JSON(newProjectCreatedResponse(&p))
}
