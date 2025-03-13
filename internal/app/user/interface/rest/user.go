package rest

import (
	"backend/internal/app/user/usecase"
	"backend/internal/domain/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserHandler struct {
	Validator   *validator.Validate
	UserUseCase usecase.UserUsecaseItf
}

func NewUserHandler(routerGroup fiber.Router, validator *validator.Validate, userUseCase usecase.UserUsecaseItf) {

	handler := &UserHandler{
		Validator:   validator,
		UserUseCase: userUseCase,
	}

	routerGroup = routerGroup.Group("/users")

	routerGroup.Post("/register", handler.RegisterUser)
	routerGroup.Post("/login", handler.Login)

}

func (h *UserHandler) RegisterUser(ctx *fiber.Ctx) error {
	var register dto.RegisterUser

	err := ctx.BodyParser(&register)
	if err != nil {
		return err
	}

	err = h.Validator.Struct(register)
	if err != nil {
		return err
	}

	err = h.UserUseCase.Register(register)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK)

}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {

	var login dto.LoginUser

	err := ctx.BodyParser(&login)
	if err != nil {
		return err
	}

	err = h.Validator.Struct(login)
	if err != nil {
		return err
	}

	token, err := h.UserUseCase.Login(login)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})

}
