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
)

func Start() error {

	config, err := env.New()
	if err != nil {
		panic(err)
	}

	database, err := mysql.New(fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
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

	return app.Listen(fmt.Sprintf(":%d", config.AppPort))

}
