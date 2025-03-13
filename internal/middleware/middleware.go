package middleware

import (
	"backend/internal/infra/jwt"
	"github.com/gofiber/fiber/v2"
)

type MiddlewareI interface {
	Authentication(ctx *fiber.Ctx) error
	Authorization(ctx *fiber.Ctx) error
}

type Middleware struct {
	jwt jwt.JWTI
}

func NewMiddleware(jwt jwt.JWTI) MiddlewareI {
	return &Middleware{
		jwt: jwt,
	}
}
