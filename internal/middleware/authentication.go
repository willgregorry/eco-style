package middleware

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func (m *Middleware) Authentication(ctx *fiber.Ctx) error {

	authToken := ctx.GetReqHeaders()["Authorization"]
	if len(authToken) < 1 {
		return ctx.Status(401).JSON(fiber.Map{
			"message": "No token",
		})
	}

	bearerToken := authToken[0]
	token := strings.Split(bearerToken, " ")

	userID, isAdmin, err := m.jwt.ValidateToken(token[1])
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	ctx.Locals("userID", userID)
	ctx.Locals("isAdmin", isAdmin)

	return ctx.Next()
}
