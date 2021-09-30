package handler

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spctr-cc/backend-bugtrack/model"
)

type accountRegisterRequest struct {
	Account struct {
		Firstname string `json:"firstname" validate:"required,min=2,max=100"`
		Lastname  string `json:"lastname" validate:"required,min=2,max=100"`
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required"`
	} `json:"account"`
}

func (r *accountRegisterRequest) bind(c *fiber.Ctx, a *model.Account, v *Validator) error {
	// validate

	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	if strings.TrimSpace(r.Account.Firstname) == "" {
		return errors.New("firstname is empty")
	}

	a.Firstname = r.Account.Firstname
	a.Lastname = r.Account.Lastname
	a.Email = r.Account.Email
	h, err := a.HashPassword(r.Account.Password)
	if err != nil {
		return err
	}
	a.Password = h

	// DEFAULT ROLE
	a.Role = model.Admin

	return nil
}

type accountLoginRequest struct {
	Account struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"account"`
}

func (r *accountLoginRequest) bind(c *fiber.Ctx, v *Validator) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type ticketCreateRequest struct {
	Ticket struct {
		Topic       string `json:"topic" validate:"required"`
		Description string `json:"description" validate:"required"`
		Priority    int    `json:"priority" validate:"required"`
		Type        int    `json:"type" validate:"required"`
		Status      int    `json:"status" validate:"required"`
	}
}

func (r *ticketCreateRequest) bind(c *fiber.Ctx, t *model.Ticket, v *Validator) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}

	t.Topic = r.Ticket.Topic
	t.Description = r.Ticket.Description
	t.Priority = model.TicketPriority(r.Ticket.Priority)
	t.Type = model.TicketType(r.Ticket.Type)
	t.Status = model.TicketStatus(r.Ticket.Status)

	return nil

}
