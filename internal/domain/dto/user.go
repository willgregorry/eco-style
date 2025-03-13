package dto

import "github.com/google/uuid"

type RegisterUser struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Role        string `json:"role" validate:"required"`
	IsAdmin     bool   `json:"is_admin"`
	GoogleID    string `json:"omitempty" validate:"omitempty,gte=1,lte=255"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserParam struct {
	Id    uuid.UUID
	Email string
}
