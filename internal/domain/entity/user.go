package entity

import (
	"backend/internal/domain/dto"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Email       string    `gorm:"type:varchar(100);not null"`
	Password    string    `gorm:"type:varchar(100);not null"`
	PhoneNumber string    `gorm:"type:varchar(100);not null"`
	Address     string    `gorm:"type:varchar(100);not null"`
	Role        string    `gorm:"type:varchar(100);not null"`
	IsAdmin     bool      `gorm:"type:boolean;not null"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	GoogleID    string    `gorm:"type:varchar(100);not null"`
}

func (u User) ParseToDTOGetUsers() dto.RequestGetUsers {
	return dto.RequestGetUsers{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		Password:    u.Password,
		PhoneNumber: u.PhoneNumber,
		Address:     u.Address,
		Role:        u.Role,
		IsAdmin:     u.IsAdmin,
	}
}

func (u User) ParseToDTOGetUsername() dto.RequestGetUsername {
	return dto.RequestGetUsername{
		Name: u.Name,
	}
}
