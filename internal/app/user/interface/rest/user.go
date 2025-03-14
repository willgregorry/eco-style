package rest

import (
	"backend/internal/app/user/usecase"
	"backend/internal/domain/dto"
	"backend/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
)

type UserHandler struct {
	Validator   *validator.Validate
	UserUseCase usecase.UserUsecaseItf
	Middleware  middleware.MiddlewareI
}

func NewUserHandler(routerGroup fiber.Router, validator *validator.Validate, userUseCase usecase.UserUsecaseItf, middleware middleware.MiddlewareI) {

	handler := &UserHandler{
		Validator:   validator,
		UserUseCase: userUseCase,
		Middleware:  middleware,
	}

	routerGroup = routerGroup.Group("/users")

	routerGroup.Get("/users", middleware.Authentication, middleware.Authorization, handler.GetAllUsers)
	routerGroup.Delete("/:id", middleware.Authentication, middleware.Authorization, handler.DeleteUser)
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

func (h UserHandler) DeleteUser(ctx *fiber.Ctx) error {

	userID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "UserID harus UUID")
	}

	err = h.UserUseCase.DeleteUser(userID)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})

}

func (h *UserHandler) GetAllUsers(ctx *fiber.Ctx) error {

	res, err := h.UserUseCase.GetAllUsers()
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": res,
	})
}
