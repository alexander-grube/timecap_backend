package handler

import "github.com/spctr-cc/backend-bugtrack/account"

type Handler struct {
	accountStore account.Store
	validator    *Validator
}

func NewHandler(as account.Store) *Handler {
	v := NewValidator()
	return &Handler{
		accountStore: as,
		validator:    v,
	}
}
