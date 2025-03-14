package repository

import (
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

type UserMySQLItf interface {
	Create(user *entity.User) error
	Get(user *entity.User, userParam dto.UserParam) error
	GetAll(user *[]entity.User) error
	Delete(user *entity.User) error
}

type UserMySQL struct {
	db *gorm.DB
}

func NewUserMySQL(db *gorm.DB) UserMySQLItf {
	return &UserMySQL{db}
}

func (r *UserMySQL) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *UserMySQL) Get(user *entity.User, userParam dto.UserParam) error {
	return r.db.First(user, userParam).Error
}

func (r *UserMySQL) GetAll(user *[]entity.User) error {
	return r.db.Find(user).Error
}

func (r *UserMySQL) Delete(user *entity.User) error {
	q := r.db.Debug().Delete(user).RowsAffected

	if q == 0 {
		return fiber.NewError(http.StatusNotFound, "User not found!")
	}

	return nil
}
