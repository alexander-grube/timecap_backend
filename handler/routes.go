package handler

import (
	"github.com/gofiber/fiber/v3"
	jwtware "github.com/spctr-cc/backend-bugtrack/middleware"
	"github.com/spctr-cc/backend-bugtrack/utils"
)

func (h *Handler) Register(r *fiber.App) {

	// root route
	r.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Is Up")
	})

	v1 := r.Group("/bugtracker")

	signingKey := jwtware.SigningKey{
		JWTAlg: "HS256",
		Key:    utils.JWTSecret,
	}

	jwtMiddleware := jwtware.New(
		jwtware.Config{
			SigningKey: signingKey,
		})

	guestUsers := v1.Group("/users")
	guestUsers.Post("/login", h.Login)
	guestUsers.Post("", h.SignUp)

	account := v1.Group("/account", jwtMiddleware)
	account.Get("", h.CurrentAccount)
	account.Get("/:id", h.GetByID)
	account.Put("", h.UpdateAccount)
	account.Delete("", h.DeleteAccount)

	ticket := v1.Group("/ticket", jwtMiddleware)
	ticket.Post("/new", h.CreateTicket)
	ticket.Get("/:id", h.GetTicketByID)
	ticket.Put("/:id", h.UpdateTicket)
	ticket.Get("", h.GetAllTickets)

	project := v1.Group("/project", jwtMiddleware)
	project.Post("/new", h.CreateProject)
	// project.Get("/:id", h.GetProjectByID)
	// project.Put("/:id", h.UpdateProject)
	// project.Get("", h.GetAllProjects)
	// project.Post("/:id/add", h.AddMemberToProject)
	// project.Post("/:id/remove", h.RemoveMemberFromProject)

	azureDevOps := v1.Group("/azure")
	azureDevOps.Post("/webhook", h.CreateAzureTicket)
	// TODO: more routes

}
