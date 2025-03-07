package mysql

import (
	"backend/internal/domain/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(entity.Product{})
	if err != nil {
		return err
	}
	return nil

}
