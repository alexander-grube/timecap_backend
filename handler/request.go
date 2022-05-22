package handler

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spctr-cc/backend-bugtrack/account"
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

type accountUpdateRequest struct {
	Account struct {
		Role      int    `json:"role"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	} `json:"account"`
}

func (r *accountUpdateRequest) bind(c *fiber.Ctx, a *model.Account, v *Validator) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	if strings.TrimSpace(r.Account.Firstname) != "" {
		a.Firstname = r.Account.Firstname
	}

	if strings.TrimSpace(r.Account.Lastname) != "" {
		a.Lastname = r.Account.Lastname
	}

	if strings.TrimSpace(r.Account.Email) != "" {
		a.Email = r.Account.Email
	}

	if strings.TrimSpace(r.Account.Password) != "" {
		h, err := a.HashPassword(r.Account.Password)
		if err != nil {
			return err
		}
		a.Password = h
	}

	if r.Account.Role != 0 {
		a.Role = model.AccountRole(r.Account.Role)
	}
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
		Priority    uint   `json:"priority" validate:"required"`
		Type        uint   `json:"type" validate:"required"`
		Status      uint   `json:"status" validate:"required"`
		Topic       string `json:"topic" validate:"required"`
		Description string `json:"description" validate:"required"`
	} `json:"ticket"`
}

func (r *ticketCreateRequest) bind(c *fiber.Ctx, t *model.Ticket, v *Validator, as account.Store) error {
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

	userID := userIDFromToken(c)

	account, err := as.GetByID(userID)

	if err != nil {
		return err
	}

	if uint(account.Role) >= uint(model.User) {
		return errors.New("user not allowed to create tickets")
	}

	t.AccountID = userID

	return nil

}

type ticketUpdateRequest struct {
	Ticket struct {
		ID          uint   `json:"id"`
		Priority    uint   `json:"priority"`
		Type        uint   `json:"type"`
		Status      uint   `json:"status"`
		Topic       string `json:"topic"`
		Description string `json:"description"`
	} `json:"ticket"`
}

func (r *ticketUpdateRequest) bind(c *fiber.Ctx, t *model.Ticket, v *Validator, as account.Store) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}

	if r.Ticket.ID == 0 {
		return errors.New("ticket id is empty")
	}

	t.ID = r.Ticket.ID
	t.Topic = r.Ticket.Topic
	t.Description = r.Ticket.Description
	t.Priority = model.TicketPriority(r.Ticket.Priority)
	t.Type = model.TicketType(r.Ticket.Type)
	t.Status = model.TicketStatus(r.Ticket.Status)

	userID := userIDFromToken(c)

	account, err := as.GetByID(userID)

	if err != nil {
		return err
	}

	if uint(account.Role) >= uint(model.User) {
		return errors.New("user not allowed to update tickets")
	}

	t.AccountID = userID

	return nil
}
