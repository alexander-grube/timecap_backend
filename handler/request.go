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
