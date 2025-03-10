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
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	GoogleID    string    `gorm:"type:varchar(100);not null"`
}

func (p User) ParseToDTO() dto.ResponseCreateUser {
	return dto.ResponseCreateUser{
		Name:        p.Name,
		Email:       p.Email,
		Password:    p.Password,
		PhoneNumber: p.PhoneNumber,
		Address:     p.Address,
		Role:        p.Role,
		GoogleID:    p.GoogleID,
	}
}
