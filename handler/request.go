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
		Mobile    string `json:"mobile" validate:"required"`
		Street    string `json:"street" validate:"required,min=2,max=100"`
		Zipcode   string `json:"zipcode" validate:"required"`
		City      string `json:"city" validate:"required"`
		Country   string `json:"country" validate:"required"`
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
	a.Mobile = r.Account.Mobile
	a.Street = r.Account.Street
	a.Zipcode = r.Account.Zipcode
	a.City = r.Account.City
	a.Country = r.Account.Country
	return nil
}
