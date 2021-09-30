package handler

import (
	"github.com/spctr-cc/backend-bugtrack/account"
	"github.com/spctr-cc/backend-bugtrack/ticket"
)

type Handler struct {
	accountStore account.Store
	ticketStore  ticket.Store
	validator    *Validator
}

func NewHandler(as account.Store, ts ticket.Store) *Handler {
	v := NewValidator()
	return &Handler{
		accountStore: as,
		ticketStore:  ts,
		validator:    v,
	}
}
