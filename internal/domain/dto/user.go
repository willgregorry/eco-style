package dto

import "github.com/google/uuid"

type RegisterUser struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"omitempty"`
	Address     string `json:"address" validate:"omitempty"`
	Role        string `json:"role" validate:"omitempty"`
	IsAdmin     bool   `json:"is_admin" validate:"omitempty"`
	GoogleID    string `json:"omitempty" validate:"omitempty,gte=1,lte=255"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RequestGetUsers struct {
	ID          uuid.UUID
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	Address     string
	Role        string
	IsAdmin     bool
	GoogleID    string
}

type UserParam struct {
	Id    uuid.UUID
	Email string
}
