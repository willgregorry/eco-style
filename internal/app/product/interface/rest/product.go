package rest

import (
	productusecase "backend/internal/app/product/usecase"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	ProductUseCase productusecase.ProductUsecaseItf
}

func NewProductHandler(routerGroup fiber.Router, productUseCase productusecase.ProductUsecaseItf) {
	handler := ProductHandler{
		ProductUseCase: productUseCase,
	}

	routerGroup = routerGroup.Group("/products")

	routerGroup.Get("/", handler.GetAllProducts)
}

func (h *ProductHandler) GetAllProducts(ctx *fiber.Ctx) error {

	res := h.ProductUseCase.Perantara()

	return ctx.JSON(fiber.Map{
		"message": res,
	})
}
