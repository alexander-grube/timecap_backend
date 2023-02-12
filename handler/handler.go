package handler

import (
	"github.com/spctr-cc/backend-bugtrack/store/account"
	"github.com/spctr-cc/backend-bugtrack/store/project"
	"github.com/spctr-cc/backend-bugtrack/store/ticket"
)

type Handler struct {
	accountStore account.Store
	ticketStore  ticket.Store
	projectStore project.Store
	validator    *Validator
}

func NewHandler(as account.Store, ts ticket.Store, ps project.Store) *Handler {
	v := NewValidator()
	return &Handler{
		accountStore: as,
		ticketStore:  ts,
		projectStore: ps,
		validator:    v,
	}
}
