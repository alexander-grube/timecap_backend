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
