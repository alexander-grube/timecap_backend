package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) Register(r *fiber.App) {
	v1 := r.Group("/api")

	// TODO: JWT LATER

	account := v1.Group("/account")
	account.Post("", h.SignUp)
}
