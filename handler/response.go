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
		Token     string            `json:"token"`
		Role      model.AccountRole `json:"role"`
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

type safeAccountResponse struct {
	Account struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Email     string `json:"email"`
	} `json:"account"`
}

func newSafeAccountResponse(a *model.Account) *safeAccountResponse {
	r := new(safeAccountResponse)
	r.Account.Firstname = a.Firstname
	r.Account.Lastname = a.Lastname
	r.Account.Email = a.Email
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

type accountDeleteResponse struct {
	Account struct {
		Firstname string            `json:"firstname"`
		Lastname  string            `json:"lastname"`
		Email     string            `json:"email"`
		Role      model.AccountRole `json:"role"`
	} `json:"account"`
}

func newAccountDeleteResponse(a *model.Account) *accountDeleteResponse {
	r := new(accountDeleteResponse)
	r.Account.Firstname = a.Firstname
	r.Account.Lastname = a.Lastname
	r.Account.Email = a.Email
	r.Account.Role = a.Role
	return r
}

type ticketResponse struct {
	Ticket struct {
		Priority    uint   `json:"priority"`
		Type        uint   `json:"type"`
		Status      uint   `json:"status"`
		Topic       string `json:"topic"`
		Description string `json:"description"`
		Account     struct {
			Role      uint   `json:"role"`
			Firstname string `json:"firstname"`
			Lastname  string `json:"lastname"`
			Email     string `json:"email"`
		} `json:"account"`
	} `json:"ticket"`
}

func newTicketResponse(t *model.Ticket, a *model.Account) *ticketResponse {
	r := new(ticketResponse)
	r.Ticket.Topic = t.Topic
	r.Ticket.Description = t.Description
	r.Ticket.Priority = uint(t.Priority)
	r.Ticket.Type = uint(t.Type)
	r.Ticket.Status = uint(t.Status)

	r.Ticket.Account.Firstname = a.Firstname
	r.Ticket.Account.Lastname = a.Lastname
	r.Ticket.Account.Email = a.Email
	r.Ticket.Account.Role = uint(a.Role)

	return r
}

type ticketOverviewResponse struct {
	Ticket []struct {
		ID          uint   `json:"id"`
		Priority    uint   `json:"priority"`
		Type        uint   `json:"type"`
		Status      uint   `json:"status"`
		Topic       string `json:"topic"`
		Description string `json:"description"`
	} `json:"ticket"`
}

func newTicketOverviewResponse(t []*model.Ticket) *ticketOverviewResponse {
	r := new(ticketOverviewResponse)
	for _, v := range t {
		r.Ticket = append(r.Ticket, struct {
			ID          uint   `json:"id"`
			Priority    uint   `json:"priority"`
			Type        uint   `json:"type"`
			Status      uint   `json:"status"`
			Topic       string `json:"topic"`
			Description string `json:"description"`
		}{
			ID:          v.ID,
			Topic:       v.Topic,
			Description: v.Description,
			Priority:    uint(v.Priority),
			Type:        uint(v.Type),
			Status:      uint(v.Status),
		})
	}
	return r
}

type ticketCreatedResponse struct {
	Ticket struct {
		Priority    uint   `json:"priority"`
		Type        uint   `json:"type"`
		Status      uint   `json:"status"`
		AccountID   uint   `json:"account_id"`
		Topic       string `json:"topic"`
		Description string `json:"description"`
	} `json:"ticket"`
}

func newTicketCreatedResponse(t *model.Ticket) *ticketCreatedResponse {
	r := new(ticketCreatedResponse)
	r.Ticket.Topic = t.Topic
	r.Ticket.Description = t.Description
	r.Ticket.Priority = uint(t.Priority)
	r.Ticket.Type = uint(t.Type)
	r.Ticket.Status = uint(t.Status)
	r.Ticket.AccountID = t.AccountID

	return r
}
