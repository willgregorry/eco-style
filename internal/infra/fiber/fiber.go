package fiber

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"time"
)

const idleTimeout = 5 * time.Second

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		IdleTimeout:  idleTimeout,
		ErrorHandler: CustomError,
	})

	return app
}

func CustomError(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Return from handler
	return ctx.Status(code).JSON(fiber.Map{
		"message": err.Error(),
	})
}
