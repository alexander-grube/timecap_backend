package handler

import (
	"github.com/spctr-cc/backend-bugtrack/model"
	"github.com/spctr-cc/backend-bugtrack/utils"
)

type accountResponse struct {
	Account struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Email     string `json:"email"`
		Token     string `json:"token"`
	} `json:"account"`
}

func newAccountResponse(a *model.Account) *accountResponse {
	r := new(accountResponse)
	r.Account.Firstname = a.Firstname
	r.Account.Lastname = a.Lastname
	r.Account.Email = a.Email
	r.Account.Token = utils.GenerateJWT(a.ID)
	return r
}
