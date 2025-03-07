package repository

import (
	"gorm.io/gorm"
)

type ProductMySQLItf interface {
	GetProducts() string
}

type ProductMySQL struct {
	db *gorm.DB
}

func NewProductMySQL(db *gorm.DB) ProductMySQLItf {
	return &ProductMySQL{db}
}

func (r ProductMySQL) GetProducts() string {
	return "it works"
}
