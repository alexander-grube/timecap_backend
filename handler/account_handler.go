package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
	jwtware "github.com/golang-jwt/jwt/v5"
	"github.com/spctr-cc/backend-bugtrack/model"
	"github.com/spctr-cc/backend-bugtrack/utils"
)

func (h *Handler) Login(c fiber.Ctx) error {
	req := &accountLoginRequest{}
	if err := req.bind(c, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	a, err := h.accountStore.GetByEmail(req.Account.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if a == nil {
		return c.Status(http.StatusForbidden).JSON(utils.AccessForbidden())
	}
	if !a.CheckPassword(req.Account.Password) {
		fmt.Printf("wrong password %v", err)
		return c.Status(http.StatusForbidden).JSON(utils.AccessForbidden())
	}
	return c.Status(http.StatusOK).JSON(newAccountResponse(a))
}

func (h *Handler) SignUp(c fiber.Ctx) error {
	var a model.Account
	req := accountRegisterRequest{}

	if err := req.bind(c, &a, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	if err := h.accountStore.Create(&a); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusCreated).JSON(newAccountResponse(&a))
}

func (h *Handler) CurrentAccount(c fiber.Ctx) error {
	a, err := h.accountStore.GetByID(userIDFromToken(c))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if a == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound())
	}
	return c.Status(http.StatusOK).JSON(newAccountResponseWithoutJWT(a))
}

func (h *Handler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	a, err := h.accountStore.GetByID(uint(idUint))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(utils.NewError(err))
	}

	return c.JSON(newSafeAccountResponse(a))
}

func (h *Handler) UpdateAccount(c fiber.Ctx) error {
	a, err := h.accountStore.GetByID(userIDFromToken(c))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	req := accountUpdateRequest{}

	if err := req.bind(c, a, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	if err := h.accountStore.Update(a); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(newAccountUpdateResponse(a))
}

func (h *Handler) DeleteAccount(c fiber.Ctx) error {
	a, err := h.accountStore.GetByID(userIDFromToken(c))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	if a == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound())
	}
	if err := h.accountStore.Delete(a); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(newAccountDeleteResponse(a))
}

func userIDFromToken(c fiber.Ctx) uint {
	var account *jwtware.Token
	l := c.Locals("user")
	if l == nil {
		return 0
	}
	account = l.(*jwtware.Token)
	id := uint(((account.Claims.(jwtware.MapClaims)["id"]).(float64)))
	return id
}
