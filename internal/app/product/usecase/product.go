package usecase

import (
	"backend/internal/app/product/repository"
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"github.com/google/uuid"
)

type ProductUsecaseItf interface {
	GetAllProducts() (*[]entity.Product, error)
	CreateProduct(request dto.RequestCreateProduct) (dto.ResponseCreateProduct, error)
}

type ProductUsecase struct {
	ProductRepository repository.ProductMySQLItf
}

func NewProductUsecase(productRepository repository.ProductMySQLItf) ProductUsecaseItf {
	return &ProductUsecase{
		ProductRepository: productRepository,
	}
}

func (u ProductUsecase) GetAllProducts() (*[]entity.Product, error) {

	products := new([]entity.Product)

	err := u.ProductRepository.GetAllProducts(products)
	if err != nil {
		return nil, err
	}

	return products, nil

}

func (u ProductUsecase) CreateProduct(request dto.RequestCreateProduct) (dto.ResponseCreateProduct, error) {

	product := entity.Product{
		ID:              uuid.New(),
		ProductName:     request.ProductName,
		ProductBrand:    request.ProductBrand,
		ProductMaterial: request.ProductMaterial,
		ProductSize:     request.ProductSize,
		Description:     request.Description,
		Price:           request.Price,
		Stock:           request.Stock,
		Category:        request.Category,
		Condition:       request.Condition,
	}

	err := u.ProductRepository.Create(product)
	if err != nil {
		return dto.ResponseCreateProduct{}, err
	}

	return product.ParseToDTO(), nil

}
