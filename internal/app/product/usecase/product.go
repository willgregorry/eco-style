package usecase

import (
	"backend/internal/app/product/repository"
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"github.com/google/uuid"
)

type ProductUsecaseItf interface {
	GetAllProducts() (*[]dto.ResponseGetProduct, error)
	CreateProduct(request dto.RequestCreateProduct) (dto.ResponseCreateProduct, error)
	GetSpecificProduct(id uuid.UUID) (dto.ResponseGetProduct, error)
	UpdateProduct(productID uuid.UUID, request dto.RequestUpdateProduct) error
	DeleteProduct(productID uuid.UUID) error
}

type ProductUsecase struct {
	ProductRepository repository.ProductMySQLItf
}

func NewProductUsecase(productRepository repository.ProductMySQLItf) ProductUsecaseItf {
	return &ProductUsecase{
		ProductRepository: productRepository,
	}
}

func (u ProductUsecase) GetAllProducts() (*[]dto.ResponseGetProduct, error) {

	products := new([]entity.Product)

	err := u.ProductRepository.GetAllProducts(products)
	if err != nil {
		return nil, err
	}

	res := make([]dto.ResponseGetProduct, len(*products))
	for i, product := range *products {
		res[i] = product.ParseToDTOGet()
	}

	return &res, nil

}

func (u ProductUsecase) CreateProduct(request dto.RequestCreateProduct) (dto.ResponseCreateProduct, error) {

	product := &entity.Product{
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

func (u ProductUsecase) GetSpecificProduct(id uuid.UUID) (dto.ResponseGetProduct, error) {

	product := &entity.Product{
		ID: id,
	}

	err := u.ProductRepository.GetSpecificProduct(product)
	if err != nil {
		return dto.ResponseGetProduct{}, err
	}

	return product.ParseToDTOGet(), err

}

func (u ProductUsecase) UpdateProduct(productID uuid.UUID, request dto.RequestUpdateProduct) error {

	product := &entity.Product{
		ID:              productID,
		ProductName:     request.ProductName,
		ProductBrand:    request.ProductBrand,
		ProductSize:     request.ProductSize,
		ProductMaterial: request.ProductMaterial,
		Description:     request.Description,
		Price:           request.Price,
		Stock:           request.Stock,
		Category:        request.Category,
		Condition:       request.Condition,
	}

	err := u.ProductRepository.Update(product)
	if err != nil {
		return err
	}

	return nil

}

func (u ProductUsecase) DeleteProduct(productID uuid.UUID) error {

	product := &entity.Product{
		ID: productID,
	}

	return u.ProductRepository.Delete(product)
}
