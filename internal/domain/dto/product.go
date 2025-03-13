package dto

type RequestCreateProduct struct {
	ProductName     string `json:"product_name" validate:"required"`
	ProductBrand    string `json:"product_brand" validate:"required"`
	ProductMaterial string `json:"product_material" validate:"required"`
	ProductSize     string `json:"product_size" validate:"required"`
	Description     string `json:"description" validate:"required"`
	Price           int64  `json:"price" validate:"required"`
	Stock           int8   `json:"stock" validate:"required"`
	Category        string `json:"category" validate:"required"`
	Condition       string `json:"condition" validate:"required"`
}

type RequestUpdateProduct struct {
	ProductName     string `json:"product_name" validate:""`
	ProductBrand    string `json:"product_brand" validate:""`
	ProductMaterial string `json:"product_material" validate:""`
	ProductSize     string `json:"product_size" validate:""`
	Description     string `json:"description" validate:""`
	Price           int64  `json:"price" validate:""`
	Stock           int8   `json:"stock" validate:""`
	Category        string `json:"category" validate:""`
	Condition       string `json:"condition" validate:""`
}

type ResponseCreateProduct struct {
	ProductName     string `json:"product_name"`
	ProductBrand    string `json:"product_brand"`
	ProductMaterial string `json:"product_material"`
	ProductSize     string `json:"product_size"`
	Description     string `json:"description"`
	Price           int64  `json:"price"`
	Stock           int8   `json:"stock"`
	Category        string `json:"category"`
	Condition       string `json:"condition"`
}

type ResponseGetProduct struct {
	ProductName     string `json:"product_name"`
	ProductBrand    string `json:"product_brand"`
	ProductMaterial string `json:"product_material"`
	ProductSize     string `json:"product_size"`
	Description     string `json:"description"`
	Price           int64  `json:"price"`
	Stock           int8   `json:"stock"`
	Category        string `json:"category"`
	Condition       string `json:"condition"`
}
