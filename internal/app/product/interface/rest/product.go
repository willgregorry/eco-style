package rest

import (
	productusecase "backend/internal/app/product/usecase"
	"backend/internal/domain/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	routerGroup.Get("/:id", handler.GetSpecificProduct)
	routerGroup.Post("/", handler.CreateProduct)
	routerGroup.Patch("/:id", handler.UpdateProduct)
	routerGroup.Delete("/:id", handler.DeleteProduct)
}

func (h ProductHandler) GetSpecificProduct(ctx *fiber.Ctx) error {

	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "ProductID harus UUID")
	}

	res, err := h.ProductUseCase.GetSpecificProduct(productID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(res)
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

func (h ProductHandler) UpdateProduct(ctx *fiber.Ctx) error {

	var request dto.RequestUpdateProduct

	err := ctx.BodyParser(&request)
	if err != nil {
		return err
	}

	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "ProductID harus UUID")
	}

	err = h.Validator.Struct(request)
	if err != nil {
		return err
	}

	err = h.ProductUseCase.UpdateProduct(productID, request)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
	})
}

func (h ProductHandler) DeleteProduct(ctx *fiber.Ctx) error {

	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "ProductID harus UUID")
	}

	err = h.ProductUseCase.DeleteProduct(productID)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Product deleted successfully",
	})

}
