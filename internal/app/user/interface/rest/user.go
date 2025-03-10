package rest

import (
	"backend/internal/app/user/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Validator   *validator.Validate
	UserUseCase usecase.UserUsecase
}

func NewUserHandler(routerGroup fiber.Router, validator *validator.Validate, userUseCase usecase.UserUsecase) {
	routerGroup = routerGroup.Group("/session")

}
