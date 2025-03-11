package repository

import (
	"backend/internal/domain/entity"
	"gorm.io/gorm"
)

type ProductMySQLItf interface {
	GetAllProducts(products *[]entity.Product) error
	Create(product entity.Product) error
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

func (r ProductMySQL) Create(product entity.Product) error {
	return r.db.Create(&product).Error
}
