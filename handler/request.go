package handler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/spctr-cc/backend-bugtrack/model"
	"github.com/spctr-cc/backend-bugtrack/store/account"
)

type accountRegisterRequest struct {
	Account struct {
		Firstname      string `json:"firstname" validate:"required,min=2,max=100"`
		Lastname       string `json:"lastname" validate:"required,min=2,max=100"`
		Email          string `json:"email" validate:"required,email"`
		Password       string `json:"password" validate:"required"`
		OrganizationID uint   `json:"organizationID" validate:"required"`
	} `json:"account"`
}

func (r *accountRegisterRequest) bind(c fiber.Ctx, a *model.Account, v *Validator) error {
	// validate

	if err := c.Bind().Body(r); err != nil {
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
	a.OrganizationID = r.Account.OrganizationID

	// DEFAULT ROLE
	a.Role = model.Admin

	return nil
}

type accountUpdateRequest struct {
	Account struct {
		Role      int    `json:"role"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	} `json:"account"`
}

func (r *accountUpdateRequest) bind(c fiber.Ctx, a *model.Account, v *Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	if strings.TrimSpace(r.Account.Firstname) != "" {
		a.Firstname = r.Account.Firstname
	}

	if strings.TrimSpace(r.Account.Lastname) != "" {
		a.Lastname = r.Account.Lastname
	}

	if strings.TrimSpace(r.Account.Email) != "" {
		a.Email = r.Account.Email
	}

	if strings.TrimSpace(r.Account.Password) != "" {
		h, err := a.HashPassword(r.Account.Password)
		if err != nil {
			return err
		}
		a.Password = h
	}

	if r.Account.Role != 0 {
		a.Role = model.AccountRole(r.Account.Role)
	}
	return nil
}

type accountLoginRequest struct {
	Account struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"account"`
}

func (r *accountLoginRequest) bind(c fiber.Ctx, v *Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type ticketCreateRequest struct {
	Ticket struct {
		Priority    uint   `json:"priority" validate:"required"`
		Type        uint   `json:"type" validate:"required"`
		Status      uint   `json:"status" validate:"required"`
		Topic       string `json:"topic" validate:"required"`
		Description string `json:"description" validate:"required"`
		ProjectID   uint   `json:"project_id" validate:"required"`
	} `json:"ticket"`
}

func (r *ticketCreateRequest) bind(c fiber.Ctx, t *model.Ticket, v *Validator, as account.Store, isServiceUser bool) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}

	t.Topic = r.Ticket.Topic
	t.Description = r.Ticket.Description
	t.Priority = model.TicketPriority(r.Ticket.Priority)
	t.Type = model.TicketType(r.Ticket.Type)
	t.Status = model.TicketStatus(r.Ticket.Status)

	var userID uint
	if !isServiceUser {
		userID = userIDFromToken(c)
	}

	if isServiceUser {
		userID = 17
	}

	account, err := as.GetByID(userID)

	if err != nil {
		return err
	}

	if uint(account.Role) >= uint(model.User) {
		return errors.New("user not allowed to create tickets")
	}

	t.AccountID = userID

	// todo for now set ProjectID to 1
	t.ProjectID = r.Ticket.ProjectID

	return nil

}

type ticketUpdateRequest struct {
	Ticket struct {
		ID          uint   `json:"id"`
		Priority    uint   `json:"priority"`
		Type        uint   `json:"type"`
		Status      uint   `json:"status"`
		Topic       string `json:"topic"`
		Description string `json:"description"`
	} `json:"ticket"`
}

func (r *ticketUpdateRequest) bind(c fiber.Ctx, t *model.Ticket, v *Validator, as account.Store) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}

	if r.Ticket.ID == 0 {
		return errors.New("ticket id is empty")
	}

	t.ID = r.Ticket.ID
	t.Topic = r.Ticket.Topic
	t.Description = r.Ticket.Description
	t.Priority = model.TicketPriority(r.Ticket.Priority)
	t.Type = model.TicketType(r.Ticket.Type)
	t.Status = model.TicketStatus(r.Ticket.Status)

	userID := userIDFromToken(c)

	account, err := as.GetByID(userID)

	if err != nil {
		return err
	}

	if uint(account.Role) >= uint(model.User) {
		return errors.New("user not allowed to update tickets")
	}

	t.AccountID = userID

	return nil
}

type projectCreateRequest struct {
	Project struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	} `json:"project"`
}

func (r *projectCreateRequest) bind(c fiber.Ctx, p *model.Project, v *Validator, as account.Store) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}

	userID := userIDFromToken(c)

	account, err := as.GetByID(userID)

	if err != nil {
		return err
	}

	if uint(account.Role) >= uint(model.User) {
		return errors.New("user not allowed to create projects")
	}

	p.Name = r.Project.Name
	p.Description = r.Project.Description
	p.Accounts = append(p.Accounts, *account)

	return nil
}

type azureCreateTicketRequest struct {
	SubscriptionID  string `json:"subscriptionId"`
	DetailedMessage struct {
		Text string `json:"text"`
	} `json:"detailedMessage"`
	Resource struct {
		ID  int    `json:"id"`
		URL string `json:"url"`
	} `json:"resource"`
}

func (r *azureCreateTicketRequest) bind(c fiber.Ctx, t *model.Ticket, v *Validator, as account.Store) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}
	if err := v.Validate(r); err != nil {
		return err
	}

	tcr := &ticketCreateRequest{}

	parseAzureText(r.DetailedMessage.Text, tcr)

	// fmt.Printf("%+v\n", tcr)

	err := tcr.bind(c, t, v, as, true)
	if err != nil {
		return err
	}

	return nil
}

func parseAzureText(text string, tcr *ticketCreateRequest) {
	// remove all \r\n from text
	text = strings.ReplaceAll(text, "\r\n", " ")
	text = strings.TrimSpace(text)
	fmt.Println(text)
	split := strings.SplitN(text, " ", 2)
	switch split[0] {
	case "Bug":
		tcr.Ticket.Type = uint(model.Bug)
	default:
		break
	}

	split = strings.SplitN(split[1], ")", 2)
	tcr.Ticket.Topic = split[0] + ")"

	split = strings.SplitN(split[1], "State: ", 2)
	tcr.Ticket.Description = "CREATED BY AZURE \n" + split[0] + split[1]

	split = strings.SplitN(split[1], "-", 2)

	split[0] = strings.TrimSpace(split[0])
	switch split[0] {
	case "New":
		tcr.Ticket.Status = uint(model.New)
	default:
		break
	}

	tcr.Ticket.Priority = uint(model.Medium)
	tcr.Ticket.ProjectID = 1
}
