package dto

type RequestCreateUser struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required,number"`
	Address     string `json:"address" validate:"required,address"`
	Role        string `json:"role" validate:"required,role"`
	GoogleID    string `json:"omitempty" validate:"omitempty,gte=1,lte=255"`
}

type ResponseCreateUser struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Role        string `json:"role"`
	GoogleID    string `json:"omitempty"`
}
