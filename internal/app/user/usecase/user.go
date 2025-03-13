package usecase

import (
	"backend/internal/app/user/repository"
	"backend/internal/domain/dto"
	"backend/internal/domain/entity"
	"backend/internal/infra/jwt"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseItf interface {
	Register(user dto.RegisterUser) error
	Login(user dto.LoginUser) (string, error)
}

type UserUsecase struct {
	userRepo repository.UserMySQLItf
	jwt      jwt.JWTI
}

func NewUserUsecase(userRepo repository.UserMySQLItf, jwt jwt.JWTI) UserUsecaseItf {
	return &UserUsecase{
		userRepo: userRepo,
		jwt:      jwt,
	}
}

func (u *UserUsecase) Register(register dto.RegisterUser) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entity.User{
		ID:          uuid.New(),
		Name:        register.Name,
		Email:       register.Email,
		PhoneNumber: register.PhoneNumber,
		Address:     register.Address,
		Role:        register.Role,
		Password:    string(hashedPassword),
		IsAdmin:     false,
	}

	err = u.userRepo.Create(user)

	return err
}

func (u *UserUsecase) Login(login dto.LoginUser) (string, error) {

	var user entity.User

	err := u.userRepo.Get(&user, dto.UserParam{
		Email: login.Email,
	})
	if err != nil {
		return "", errors.New("email or password is invalid")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return "", errors.New("email or password is invalid")
	}

	token, err := u.jwt.GenerateToken(user.ID, user.IsAdmin)
	if err != nil {
		return "", err
	}

	return token, nil

}
