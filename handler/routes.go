package handler

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/spctr-cc/backend-bugtrack/utils"
)

func (h *Handler) Register(r *fiber.App) {
	v1 := r.Group("/api")
	jwtMiddleware := jwtware.New(
		jwtware.Config{
			SigningKey: utils.JWTSecret,
		})

	guestUsers := v1.Group("/users")
	guestUsers.Post("/login", h.Login)
	guestUsers.Post("", h.SignUp)
	account := v1.Group("/account", jwtMiddleware)
	account.Get("", h.CurrentAccount)
	ticket := v1.Group("/ticket", jwtMiddleware)
	ticket.Post("/new", h.CreateTicket)
}
