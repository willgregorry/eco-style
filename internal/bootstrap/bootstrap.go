package bootstrap

import (
	producthandler "backend/internal/app/product/interface/rest"
	productrepository "backend/internal/app/product/repository"
	productusecase "backend/internal/app/product/usecase"
	"backend/internal/infra/env"
	"backend/internal/infra/mysql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
)

func Start() error {

	config, err := env.New()
	if err != nil {
		config = &env.Env{
			AppPort:    os.Getenv("APP_PORT"),
			DBUsername: os.Getenv("DB_USERNAME"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBHost:     os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
			DBName:     os.Getenv("DB_NAME"),
		}
	}

	database, err := mysql.New(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	))

	err = mysql.Migrate(database)
	if err != nil {
		panic(err)
	}

	val := validator.New()

	app := fiber.New()

	v1 := app.Group("/api/v1")

	productRepository := productrepository.NewProductMySQL(database)
	productUseCase := productusecase.NewProductUsecase(productRepository)
	producthandler.NewProductHandler(v1, val, productUseCase)

	appPort, _ := strconv.Atoi(config.AppPort)
	return app.Listen(fmt.Sprintf(":%d", appPort))

}
