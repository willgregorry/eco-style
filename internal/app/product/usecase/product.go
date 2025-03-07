package usecase

import "backend/internal/app/product/repository"

type ProductUsecaseItf interface {
	Perantara() string
}

type ProductUsecase struct {
	ProductRepository repository.ProductMySQLItf
}

func NewProductUsecase(productRepository repository.ProductMySQLItf) ProductUsecaseItf {
	return &ProductUsecase{
		ProductRepository: productRepository,
	}
}

func (u ProductUsecase) Perantara() string {
	return u.ProductRepository.GetProducts()
}
