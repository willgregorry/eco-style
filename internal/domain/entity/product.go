package entity

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID              uuid.UUID `gorm:"type:char(36);primary_key"`
	ProductName     string    `gorm:"type:varchar(100);not null"`
	ProductBrand    string    `gorm:"type:varchar(100);not null"`
	ProductMaterial string    `gorm:"type:varchar(100);not null"`
	ProductSize     string    `gorm:"type:varchar(100);not null"`
	Description     string    `gorm:"type:text"`
	Price           int64     `gorm:"type:bigint;not null"`
	Stock           int8      `gorm:"type:smallint;not null"`
	Category        string    `gorm:"type:char(36);not null"`
	Condition       string    `gorm:"type:text;not null"`
	PhotoUrls       string    `gorm:"type:text"`
	CreatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
