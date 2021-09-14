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
		Mobile    string `json:"mobile"`
		Street    string `json:"street"`
		Zipcode   string `json:"zipcode"`
		City      string `json:"city"`
		Country   string `json:"country"`
		Token     string `json:"token"`
	} `json:"account"`
}

func newAccountResponse(a *model.Account) *accountResponse {
	r := new(accountResponse)
	r.Account.Firstname = a.Firstname
	r.Account.Lastname = a.Lastname
	r.Account.Email = a.Email
	r.Account.Mobile = a.Mobile
	r.Account.Street = a.Street
	r.Account.Zipcode = a.Zipcode
	r.Account.City = a.City
	r.Account.Country = a.Country
	r.Account.Token = utils.GenerateJWT(a.ID)
	return r
}
