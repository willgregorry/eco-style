package repository

import (
	"backend/internal/domain/entity"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

type ProductMySQLItf interface {
	GetAllProducts(products *[]entity.Product) error
	GetSpecificProduct(products *entity.Product) error
	Create(product *entity.Product) error
	Update(product *entity.Product) error
	Delete(product *entity.Product) error
}

type ProductMySQL struct {
	db *gorm.DB
}

func NewProductMySQL(db *gorm.DB) ProductMySQLItf {
	return &ProductMySQL{db}
}

func (r ProductMySQL) GetAllProducts(products *[]entity.Product) error {
	return r.db.Find(products).Error
}

func (r ProductMySQL) GetSpecificProduct(products *entity.Product) error {
	return r.db.First(products).Error
}

func (r ProductMySQL) Update(product *entity.Product) error {
	return r.db.Debug().Updates(product).Error
}

func (r ProductMySQL) Create(product *entity.Product) error {
	return r.db.Create(&product).Error
}

func (r ProductMySQL) Delete(product *entity.Product) error {
	q := r.db.Debug().Delete(product).RowsAffected

	if q == 0 {
		return fiber.NewError(http.StatusNotFound, "Product not found!")
	}

	return nil
}
