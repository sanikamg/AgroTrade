//go:build wireinject
// +build wireinject

package di

import (
	http "golang_project_ecommerce/pkg/api"
	"golang_project_ecommerce/pkg/api/handler"
	"golang_project_ecommerce/pkg/config"
	"golang_project_ecommerce/pkg/db"
	"golang_project_ecommerce/pkg/repository"
	"golang_project_ecommerce/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabase,
		//repository
		repository.NewUserRepository, repository.NewadminRepository, repository.NewProductRepository,

		//usecase
		usecase.NewUserUsecase, usecase.NewadminUsecase, usecase.NewProductUsecase,

		//handler
		handler.NewUserhandler, handler.NewAdminHandler, handler.NewProductHandler,

		http.NewServerHTTP,
	)
	return &http.ServerHTTP{}, nil
}
