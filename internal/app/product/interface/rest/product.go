package rest

import (
	productusecase "backend/internal/app/product/usecase"
	"backend/internal/domain/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ProductHandler struct {
	Validator      *validator.Validate
	ProductUseCase productusecase.ProductUsecaseItf
}

func NewProductHandler(routerGroup fiber.Router, validator *validator.Validate, productUseCase productusecase.ProductUsecaseItf) {
	handler := ProductHandler{
		Validator:      validator,
		ProductUseCase: productUseCase,
	}

	routerGroup = routerGroup.Group("/products")

	routerGroup.Get("/", handler.GetAllProducts)
	routerGroup.Post("/", handler.CreateProduct)

}

func (h *ProductHandler) GetAllProducts(ctx *fiber.Ctx) error {

	res, err := h.ProductUseCase.GetAllProducts()
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": res,
	})
}

func (h *ProductHandler) CreateProduct(ctx *fiber.Ctx) error {

	var request dto.RequestCreateProduct

	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	// validate
	err = h.Validator.Struct(request)
	if err != nil {
		return err
	}

	res, err := h.ProductUseCase.CreateProduct(request)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Product created successfully",
		"payload": res,
	})
}
