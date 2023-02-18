package handler

import (
	"github.com/spctr-cc/backend-bugtrack/store/account"
	"github.com/spctr-cc/backend-bugtrack/store/azure"
	"github.com/spctr-cc/backend-bugtrack/store/project"
	"github.com/spctr-cc/backend-bugtrack/store/ticket"
)

type Handler struct {
	accountStore account.Store
	ticketStore  ticket.Store
	projectStore project.Store
	azureStore   azure.Store
	validator    *Validator
}

func NewHandler(as account.Store, ts ticket.Store, ps project.Store, azs azure.Store) *Handler {
	v := NewValidator()
	return &Handler{
		accountStore: as,
		ticketStore:  ts,
		projectStore: ps,
		azureStore:   azs,
		validator:    v,
	}
}
