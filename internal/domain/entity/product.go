package entity

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key"`
	Title       string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	Price       int64     `gorm:"type:bigint;not null"`
	Stock       int8      `gorm:"type:smallint;not null"`
	PhotoUrls   string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"type:timestamp;autoUpdateTime"`
}
