package middleware

import "github.com/gofiber/fiber/v2"

func (m *Middleware) Authorization(ctx *fiber.Ctx) error {

	isAdmin := ctx.Locals("isAdmin")
	if isAdmin == false {
		return ctx.Status(401).JSON(fiber.Map{
			"message": "Admin Only",
		})
	}

	return ctx.Next()
}

func (m *Middleware) SellerAuthorization(ctx *fiber.Ctx) error {

	role := ctx.Locals("role")
	if role == "user" {
		return ctx.Status(401).JSON(fiber.Map{
			"message": "Seller/Admin Only",
		})
	}
	if role == "admin" || role == "mitra" {
		return ctx.Next()
	}

	return ctx.Next()

}
