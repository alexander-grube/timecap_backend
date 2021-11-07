package handler

import (
	"github.com/spctr-cc/backend-bugtrack/model"
	"github.com/spctr-cc/backend-bugtrack/utils"
)

type accountResponse struct {
	Account struct {
		Firstname string            `json:"firstname"`
		Lastname  string            `json:"lastname"`
		Email     string            `json:"email"`
		Role      model.AccountRole `json:"role"`
		Token     string            `json:"token"`
	} `json:"account"`
}

func newAccountResponse(a *model.Account) *accountResponse {
	r := new(accountResponse)
	r.Account.Firstname = a.Firstname
	r.Account.Lastname = a.Lastname
	r.Account.Email = a.Email
	r.Account.Role = a.Role
	r.Account.Token = utils.GenerateJWT(a.ID)
	return r
}

type accountUpdateResponse struct {
	Account struct {
		Firstname string            `json:"firstname"`
		Lastname  string            `json:"lastname"`
		Email     string            `json:"email"`
		Role      model.AccountRole `json:"role"`
	} `json:"account"`
}

func newAccountUpdateResponse(a *model.Account) *accountUpdateResponse {
	r := new(accountUpdateResponse)
	r.Account.Firstname = a.Firstname
	r.Account.Lastname = a.Lastname
	r.Account.Email = a.Email
	r.Account.Role = a.Role
	return r
}

type ticketResponse struct {
	Ticket struct {
		Topic       string `json:"topic"`
		Description string `json:"description"`
		Priority    uint   `json:"priority"`
		Type        uint   `json:"type"`
		Status      uint   `json:"status"`
		AccountID   int    `json:"accountID"`
	} `json:"ticket"`
}

func newTicketResponse(t *model.Ticket) *ticketResponse {
	r := new(ticketResponse)
	r.Ticket.Topic = t.Topic
	r.Ticket.Description = t.Description
	r.Ticket.Priority = uint(t.Priority)
	r.Ticket.Type = uint(t.Type)
	r.Ticket.Status = uint(t.Status)
	r.Ticket.AccountID = t.AccountID
	return r
}
