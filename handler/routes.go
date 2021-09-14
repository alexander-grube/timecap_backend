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
			AuthScheme: "Token",
		})

	guestUsers := v1.Group("/users")
	guestUsers.Post("/login", h.Login)

	account := v1.Group("/account", jwtMiddleware)
	account.Post("", h.SignUp)
	account.Get("", h.CurrentAccount)
}
