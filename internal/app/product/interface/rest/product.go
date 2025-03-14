package rest

import (
	productusecase "backend/internal/app/product/usecase"
	"backend/internal/domain/dto"
	"backend/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
)

type ProductHandler struct {
	Validator      *validator.Validate
	ProductUseCase productusecase.ProductUsecaseItf
	Middleware     middleware.MiddlewareI
}

func NewProductHandler(routerGroup fiber.Router, validator *validator.Validate, productUseCase productusecase.ProductUsecaseItf, middleware middleware.MiddlewareI) {
	handler := ProductHandler{
		Validator:      validator,
		ProductUseCase: productUseCase,
		Middleware:     middleware,
	}

	routerGroup = routerGroup.Group("/products")

	routerGroup.Get("/", middleware.Authentication, handler.GetAllProducts)
	routerGroup.Get("/:product_name", middleware.Authentication, handler.GetSpecificProduct)
	routerGroup.Post("/", middleware.Authentication, middleware.SellerAuthorization, handler.CreateProduct)
	routerGroup.Patch("/:id", middleware.Authentication, middleware.SellerAuthorization, handler.UpdateProduct)
	routerGroup.Delete("/:id", middleware.Authentication, middleware.SellerAuthorization, handler.DeleteProduct)

}

func (h ProductHandler) GetSpecificProduct(ctx *fiber.Ctx) error {

	parameter := ctx.Params("product_name")
	if parameter == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Product tidak ditemukan")
	}

	res, err := h.ProductUseCase.GetSpecificProduct(parameter)
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
